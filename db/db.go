package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ! keys for pasto ....
var TOKEN_SYMMETRIC_KEY = "12345678912345678912345678912345"
var ACCESS_TOKEN_DURATION = "15m"
var REFRESH_TOKEN_DURATION = "5m"

// connection to db
func ConnectToDb() *pgxpool.Pool {

	var pool *pgxpool.Pool
	var err error

	env := os.Getenv("ENVIRONMENT")

	switch env {
	case "DEV":
		pool, err = pgxpool.New(context.Background(), os.Getenv("DB_ADDRESS"))
		fmt.Println("Running in developement environment")

	default:
		fmt.Println("Unknown environment")
	}

	if err != nil {
		logFatal(err)
	}

	fmt.Println("Successfully connected!")
	return pool
}

// migrattion cli configuration...
func Migrate() error {
	_, b, _, _ := runtime.Caller(0)

	var migrationPath string

	switch runtime.GOOS {
	case "windows":
		migrationPath = fmt.Sprintf("file://%s/migrations", path.Dir(b))
		fmt.Println("Running on Windows")
	case "darwin":
		migrationPath = fmt.Sprintf("file:///%s/migrations", path.Dir(b))
		fmt.Println("Running on macOS")
	case "linux":
		migrationPath = fmt.Sprintf("file:///%s/migrations", path.Dir(b))
		fmt.Println("Running on Linux")
	default:
		migrationPath = fmt.Sprintf("file:///%s/migrations", path.Dir(b))
		fmt.Printf("Running on an unknown operating system: %s\n", runtime.GOOS)
	}

	log.Println("testing env", os.Getenv("DB_ADDRESS"))
	m, err := migrate.New(migrationPath, os.Getenv("DB_ADDRESS"))
	if err != nil {
		return fmt.Errorf("error create the migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error migrate up: %v", err)
	}

	log.Println("migration done")
	return nil
}
