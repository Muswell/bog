// +build test

package database

import (
	"log"
	"os"
)

func init() {
	dbName = os.Getenv("LLC_TEST_DB_NAME")
	dbHost = os.Getenv("LLC_TEST_DB_HOST")
	dbUser = os.Getenv("LLC_TEST_DB_USER")
	dbPass = os.Getenv("LLC_TEST_DB_PASSWORD")

	log.Println("Running database - test init")
}
