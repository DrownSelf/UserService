package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func Up(db *sql.DB) error {
	goose.SetDialect("postgres")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}

func Down(db *sql.DB) error {
	goose.SetDialect("postgres")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Down(db, "."); err != nil {
		return err
	}
	return nil
}
