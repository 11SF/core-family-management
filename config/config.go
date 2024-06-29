package config

import (
	"log/slog"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Server  appConfig
}

type appConfig struct {
	Port string
}

func InitConfig() *Config {

	config := new(Config)

	slog.Info("[env] start loading env")
	err := godotenv.Load("./configs/.env")
	if err != nil {
		slog.Error("[env] unable to load .env file", "error", err)
	}
	err = env.Parse(config)
	if err != nil {
		slog.Error("[env] unable to parse ennvironment variables", "error", err)
		panic(0)
	}
	slog.Info("[env] loading env complete")

	return config
}
