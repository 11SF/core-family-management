package config

import (
	"log/slog"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string `env:"APP_NAME"`
	Server    appConfig
	Database  databaseConfig
	Redis     redisConfig
	RedisKeys RedisKeys
}

type appConfig struct {
	Port string `env:"APP_PORT"`
}

type databaseConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
}

type redisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}

type RedisKeys struct {
	Family     string `env:"REDIS_KEY_FAMILY"`
	FamilyList string `env:"REDIS_KEY_FAMILY_LIST"`
}

func InitConfig() *Config {

	config := new(Config)

	slog.Info("[env] start loading env")
	err := godotenv.Load("./config/.env")
	if err != nil {
		slog.Error("[env] unable to load .env file", "error", err)
		panic(0)
	}
	err = env.Parse(config)
	if err != nil {
		slog.Error("[env] unable to parse ennvironment variables", "error", err)
		panic(0)
	}
	slog.Info("[env] loading env complete")

	return config
}
