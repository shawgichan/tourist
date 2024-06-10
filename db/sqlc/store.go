package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}

type SQLStore struct {
	ConnPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  New(connPool),
		ConnPool: connPool,
	}

}

func ExecuteTx(c context.Context, pool *pgxpool.Pool, s Store, fn func(q Store) error) error {
	tx, err := pool.Begin(c)
	if err != nil {
		return err
	}

	q := New(tx)
	s = q
	err = fn(s)
	if err != nil {
		if rbErr := tx.Rollback(c); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(c)
}
func New2(db DBTX) Querier {
	return &Queries{db: db}
}

func ExecTx2(c context.Context, pool *pgxpool.Pool, fn func(Store) error) error {
	tx, err := pool.Begin(c)
	if err != nil {
		return err
	}

	q := New2(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(c); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(c)
}
