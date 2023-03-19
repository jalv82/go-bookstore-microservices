package tests

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func runMigrations(db *sql.DB, scriptPath string) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("database driver could not be get")
	}

	migrate, err := migrate.NewWithDatabaseInstance("file://"+scriptPath, "postgres", driver)
	if err != nil {
		log.Fatal().Err(err).Msg("sql scripts could not be read")
	}

	err = migrate.Up()
	if err != nil {
		log.Fatal().Err(err).Msg("sql scripts could not be run")
	}
}
