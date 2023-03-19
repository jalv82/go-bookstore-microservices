package tests

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"bookstore/internal/commons"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const errMessage = "database container could not be run"

// PostgresContainer represents the postgres container type used in the module
type PostgresContainer struct {
	testcontainers.Container
}

type postgresContainerOption func(containerRequest *testcontainers.ContainerRequest)

func runPostgresContainer(ctx context.Context, databaseConfig commons.DatabaseConfig) (container *PostgresContainer, containerPort string, db *sql.DB) {
	port, err := nat.NewPort("tcp", databaseConfig.Port)
	if err != nil {
		log.Fatal().Err(err).Msg(errMessage)
	}

	container, err = startContainer(ctx,
		withPort(port.Port()),
		withInitialDatabase(databaseConfig.User, databaseConfig.Password, databaseConfig.Schema),
		withWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal().Err(err).Msg(errMessage)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg(errMessage)
	}

	containerExternalPort, err := container.MappedPort(ctx, port)
	if err != nil {
		log.Fatal().Err(err).Msg(errMessage)
	}

	containerPort = containerExternalPort.Port()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, containerPort, databaseConfig.User, databaseConfig.Password, databaseConfig.Schema)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal().Err(err).Msg(errMessage)
	}

	return
}

// startContainer creates an instance of the postgres container type
func startContainer(ctx context.Context, opts ...postgresContainerOption) (*PostgresContainer, error) {
	containerRequest := testcontainers.ContainerRequest{
		Image:        "postgres:15.1-alpine",
		Env:          map[string]string{},
		ExposedPorts: []string{},
		Cmd:          []string{"postgres", "-c", "fsync=off"},
	}

	for _, opt := range opts {
		opt(&containerRequest)
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderDocker,
		ContainerRequest: containerRequest,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{Container: container}, nil
}

func withPort(port string) func(containerRequest *testcontainers.ContainerRequest) {
	return func(containerRequest *testcontainers.ContainerRequest) {
		containerRequest.ExposedPorts = append(containerRequest.ExposedPorts, port)
	}
}

func withInitialDatabase(user string, password string, schema string) func(containerRequest *testcontainers.ContainerRequest) {
	return func(containerRequest *testcontainers.ContainerRequest) {
		containerRequest.Env["POSTGRES_USER"] = user
		containerRequest.Env["POSTGRES_PASSWORD"] = password
		containerRequest.Env["POSTGRES_DB"] = schema
	}
}

func withWaitStrategy(strategies ...wait.Strategy) func(containerRequest *testcontainers.ContainerRequest) {
	return func(containerRequest *testcontainers.ContainerRequest) {
		containerRequest.WaitingFor = wait.ForAll(strategies...).WithDeadline(1 * time.Minute)
	}
}
