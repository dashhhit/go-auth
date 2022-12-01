package migrations

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func SetupGoose(connStr string) {
	mdb, _ := sql.Open("postgres", connStr)

	err := mdb.Ping()
	if err != nil {
		log.Fatalf("open db error: %v\n", err)
	}

	err = goose.Up(mdb, "migrations")
	if err != nil {
		log.Fatalf("goose up error %v\n", err)
	}
}
