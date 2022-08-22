package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(up, down)
}

func up(tx *sql.Tx) error {
	query := `create table users(
    	"id" serial primary key,
    	"name" text,
    	"phoneNumber" text,
    	"email" text,
    	"password" text,
    	"rating" integer,
    	"created_at" date,
    	"updated_at" date,
	unique("phoneNumber", "email"));`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func down(tx *sql.Tx) error {
	query := `drop table users;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
