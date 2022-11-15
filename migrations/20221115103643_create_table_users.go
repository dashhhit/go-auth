package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTableUsers, downCreateTableUsers)
}

func upCreateTableUsers(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    gender varchar(1) NOT NULL,
    email varchar(50) NOT NULL,
    password varchar NOT NULL
);`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateTableUsers(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return err
	}
	return nil
}
