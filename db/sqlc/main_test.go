package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var testingStore Store

func init() {
	gotenv.Load()
}

func TestMain(m *testing.M) {

	pool, err := pgxpool.New(context.Background(), "postgresql://shawgi:Shawgi%40123@46.101.204.174:5432/shawgi?sslmode=disable")
	if err != nil {
		log.Fatal("erro while connecting to db ", err)
	}

	testingStore = New(pool)
	os.Exit(m.Run())
}
