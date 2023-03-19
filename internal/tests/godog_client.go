package tests

import (
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func TestSuite(t *testing.T, scenarioContext func(feature *godog.ScenarioContext), pathFeature string) (testSuite godog.TestSuite) {
	testSuite = godog.TestSuite{
		ScenarioInitializer: scenarioContext,
		Options: &godog.Options{
			Concurrency: 1,
			Format:      "pretty",
			Paths:       []string{pathFeature},
			Randomize:   time.Now().UTC().UnixNano(), // randomize scenario execution order
			TestingT:    t,                           // Testing instance that will run subtests.
		},
	}

	return
}
