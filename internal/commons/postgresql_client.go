package commons

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Image    string
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Schema   string
}

func ReadDatabaseConfig(path string) (databaseConfig *DatabaseConfig) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("the database configuration file could not be read")
	}

	databaseConfig = &DatabaseConfig{
		Image:    viper.GetString("database.image"),
		Driver:   viper.GetString("database.driver"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Schema:   viper.GetString("database.schema"),
	}

	return
}

func NewPostgreSQLClient(config *DatabaseConfig) (db *gorm.DB) {
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.Driver, config.User, config.Password, config.Host, config.Port, config.Schema)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("postgresql-client could not be created")
	}

	return
}

func NewPostgreSQLClientMock() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal().Err(err).Msgf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDb}))
	if err != nil {
		log.Fatal().Err(err).Msg("postgresql-client mock could not be created")
	}

	return db, mock
}
