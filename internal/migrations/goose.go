package migrations

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func SetupGoose(connStr string) {
	fmt.Println(connStr)
	mdb, _ := sql.Open("postgres", connStr)

	err := mdb.Ping()
	if err != nil {
		log.Fatalf("open db error: %v\n", err)
	}

	err = goose.Up(mdb, "/")
	if err != nil {
		log.Fatalf("goose up error %v\n", err)
	}
}
