package db

import (
	"database/sql"
	"github.com/kamalm96/backend/db/utils"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load test config file:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
