// +build postgres_service

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	mig "github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"log"
	"os"
)

func init() {
	log.Println("Running postgres_service DB init")
	New = func() (*sql.DB, error) {
		return sql.Open("postgres", "postgres://"+dbHost+"/"+dbName+"?user="+dbUser+"&password="+dbPass)
	}

	Up = func() error {
		log.Println("Running Up migration")
		return migrate("up")
	}

	Down = func() error {
		log.Println("Running Down migration")
		return migrate("down")
	}
}

func migrate(dir string) error {
	db, err := New()
	if err != nil {
		return err
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	defer driver.Close()

	file := os.Getenv("GOPATH") + "/src/github.com/muswell/bog/migrations/"

	m, err := mig.NewWithDatabaseInstance(
		"file://"+file,
		"postgres", driver)

	if err != nil {
		return err
	}

	switch dir {
	case "up":
		err = m.Up()
		break
	case "down":
		err = m.Down()
		break
	default:
		return fmt.Errorf("migrate dir must be up or down")
		break
	}

	if err == mig.ErrNoChange {
		log.Println("No migrations to run.")

		return nil
	}

	return err
}
