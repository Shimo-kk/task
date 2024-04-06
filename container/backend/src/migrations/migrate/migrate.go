package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func UpMigration(sourceDir string, database string) error {
	sourceURL := fmt.Sprintf("file://%s", sourceDir)
	databaseURL := fmt.Sprintf("postgres://%s?sslmode=disable", database)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func DownMigration(sourceDir string, database string) error {
	sourceURL := fmt.Sprintf("file://%s", sourceDir)
	databaseURL := fmt.Sprintf("postgres://%s?sslmode=disable", database)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
