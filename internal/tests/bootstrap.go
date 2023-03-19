package tests

import (
	"context"
	"database/sql"

	"bookstore/internal/commons"
	"github.com/rs/zerolog/log"
)

type TestEnvironment struct {
	DbConfig    *commons.DatabaseConfig
	DbContainer *PostgresContainer
	Db          *sql.DB
}

var testEnvironment *TestEnvironment

func GetTestEnvironment() *TestEnvironment {
	return testEnvironment
}
func SetUpEnvironment(ctx context.Context, databaseConfig commons.DatabaseConfig, migrateScriptPath string) {
	container, containerPort, db := runPostgresContainer(ctx, databaseConfig)

	runMigrations(db, migrateScriptPath)

	databaseConfig.Port = containerPort

	testEnvironment = &TestEnvironment{
		DbConfig:    &databaseConfig,
		DbContainer: container,
		Db:          db,
	}
}

func TearDownEnvironment(ctx context.Context) {
	err := testEnvironment.Db.Close()
	if err != nil {
		log.Fatal().Err(err).Msg("database could not be closed")
	}

	err = testEnvironment.DbContainer.Terminate(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("database container could not be closed")
	}
}
