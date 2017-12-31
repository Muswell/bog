// A package for data access
package database

import (
	"database/sql"
)

type dbFactory func() (*sql.DB, error)

var (
	New    dbFactory
	Up     func() error
	Down   func() error
	dbName string
	dbHost string
	dbUser string
	dbPass string
)
