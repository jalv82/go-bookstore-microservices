package acceptance_tests

import (
	"context"
	"os"
	"testing"

	"bookstore/internal/commons"
	"bookstore/internal/tests"
)

const (
	pathBookFeature    = "features/book.feature"
	pathConfigFile     = "config-tests.yaml"
	pathMigratorScript = "../scripts/migrator"

	errorMessage = "non-zero status returned, failed to run feature tests"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	databaseConfig := commons.ReadDatabaseConfig(pathConfigFile)

	tests.SetUpEnvironment(ctx, *databaseConfig, pathMigratorScript)
	exitCode := m.Run()
	tests.TearDownEnvironment(ctx)

	os.Exit(exitCode)
}

func TestBookFeature(t *testing.T) {
	suite := tests.TestSuite(t, InitializeBookScenario, pathBookFeature)

	if suite.Run() != 0 {
		t.Fatal(errorMessage)
	}
}
