package http_server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type HttpServerConfig struct {
	Port string
}

func ReadHttpServerConfig(path string) HttpServerConfig {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("the http server configuration file could not be read")
	}

	return HttpServerConfig{
		Port: viper.GetString("http-server.port"),
	}
}

type HttpServer struct {
	Config HttpServerConfig
	Server *echo.Echo
}

type HttpServerCtx = echo.Context

// NewHttpServer return new http instance
func NewHttpServer(config HttpServerConfig) *HttpServer {
	return &HttpServer{
		Server: echo.New(),
		Config: config,
	}
}

// Up run the http server
func (hs *HttpServer) Up() error {
	err := hs.Server.Start(":" + hs.Config.Port)
	if err != nil {
		hs.Server.Logger.Fatal(err)

		return err
	}

	return nil
}
