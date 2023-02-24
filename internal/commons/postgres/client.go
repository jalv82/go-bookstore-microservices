package postgres

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Schema   string
}

func NewSqlClient(config *DatabaseConfig) (db *gorm.DB, err error) {
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.Driver, config.User, config.Password, config.Host, config.Port, config.Schema)

	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	return
}

func NewMockSqlClient() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal().Err(err).Msgf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDb}))
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	return db, mock, err
}
