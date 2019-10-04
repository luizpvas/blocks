package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq" // We need to pull postgres to setup testing
)

// DB is a global shared database connection pool.
var DB *sql.DB

// SetConnection updates the global shared database connection pool
// to the given value. This function should be called once from the entrypoint.
func SetConnection(db *sql.DB) {
	DB = db
}

// SetTestingConnection updates the global shared connection using txdb for testing.
func SetTestingConnection(t *testing.T) func() {
	txdb.Register("txdb", "postgres", "user=blocks password=blocks dbname=blocks_test")
	db, err := sql.Open("txdb", "testing")
	if err != nil {
		t.Fatalf("could not connect to postgres: %v", err)
	}
	SetConnection(db)
	err = RunPendingMigrations()
	if err != nil {
		t.Fatalf("could not run migrations: %v", err)
	}
	return func() {
		db.Close()
	}
}

// RunPendingMigrations reads
func RunPendingMigrations() error {
	err := ensureMigrationsTableExists()
	if err != nil {
		return fmt.Errorf("could not create migrations table: %v", err)
	}

	migrations, err := executedMigrations()
	if err != nil {
		return fmt.Errorf("could not fetch executed migrations: %v", err)
	}

	all, err := readAllMigrations()
	if err != nil {
		return fmt.Errorf("could not read migrations from directory: %v", err)
	}

	for _, migration := range all {
		executed := false
		for _, alreadyRan := range migrations {
			if migration == alreadyRan {
				executed = true
				break
			}
		}

		if !executed {
			if err = runMigration(migration); err != nil {
				return fmt.Errorf("could not run pending migration: %v", err)
			}
		}
	}

	return nil
}

func readAllMigrations() ([]string, error) {
	files, err := ioutil.ReadDir("../migrations")
	if err != nil {
		return nil, err
	}

	var fnames []string
	for _, f := range files {
		if strings.Contains(f.Name(), "up") {
			fnames = append(fnames, f.Name())
		}
	}
	return fnames, nil
}

func executedMigrations() ([]string, error) {
	rows, err := DB.Query("SELECT migration_file FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return []string{}, nil
}

func ensureMigrationsTableExists() error {
	query := `
        create table if not exists migrations (
            migration_file text,
            run_at         timestamp with time zone
        );
    `

	_, err := DB.Exec(query)
	return err
}

func runMigration(file string) error {
	data, err := ioutil.ReadFile("../migrations/" + file)
	if err != nil {
		return err
	}

	statements := strings.Split(string(data), ";")
	for _, st := range statements {
		if _, err = DB.Exec(st); err != nil {
			return err
		}
	}

	return nil
}
