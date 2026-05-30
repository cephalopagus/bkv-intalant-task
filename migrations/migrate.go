package migrate

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed sql
var migrationsFS embed.FS

func Up(db *sql.DB) error {
	goose.SetBaseFS(migrationsFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}
	if err := goose.Up(db, "sql"); err != nil {
		return fmt.Errorf("goose.Up: %w", err)
	}
	return nil
}
