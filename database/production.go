// +build production

package database

import "os"

func init() {
	dbName = os.Getenv("LLC_DB_NAME")
	dbHost = os.Getenv("LLC_DB_HOST")
	dbUser = os.Getenv("LLC_DB_USER")
	dbPass = os.Getenv("LLC_DB_PASSWORD")
}
