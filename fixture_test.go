package bog

import (
	"github.com/muswell/bog/database"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := database.Up()
	if err != nil {
		log.Fatalf("Cannot initialize database: %v", err)
	}

	runTests := m.Run()

	err = database.Down()
	if err != nil {
		log.Fatalf("Cannot teardown database: %v", err)
	}
	os.Exit(runTests)
}
