package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
)

func TestMigrationsTableIsCreated(t *testing.T) {
	txdb.Register("txdb", "postgres", "user=blocks password=blocks dbname=blocks_test")
	db, err := sql.Open("txdb", "testing")
	if err != nil {
		t.Fatalf("could not connect to postgres: %v", err)
	}
	defer db.Close()
	SetConnection(db)

	err = RunPendingMigrations()
	if err != nil {
		t.Fatalf("could not rung pending migrations: %v", err)
	}
}
