package acceptance_tests

import (
	"testing"

	"github.com/cucumber/godog"
)

const (
	authorFeature = "features/author.feature"
	bookFeature   = "features/book.feature"

	errorMessage = "non-zero status returned, failed to run feature tests"
)

func commonSuite(t *testing.T, scenarioContext func(feature *godog.ScenarioContext), pathFeature string) godog.TestSuite {
	suite := godog.TestSuite{
		ScenarioInitializer: scenarioContext,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{pathFeature},
			TestingT: t, // Testing instance that will run subtests.
		},
	}
	return suite
}

func TestAuthorFeature(t *testing.T) {
	suite := commonSuite(t, InitializeAuthorScenario, authorFeature)

	if suite.Run() != 0 {
		t.Fatal(errorMessage)
	}
}

func TestBookFeature(t *testing.T) {
	suite := commonSuite(t, InitializeBookScenario, bookFeature)

	if suite.Run() != 0 {
		t.Fatal(errorMessage)
	}
}
