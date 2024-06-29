package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/11SF/core-family-management/config"
	routes "github.com/11SF/core-family-management/pkg"
	"github.com/11SF/go-common/database"
	"github.com/11SF/go-common/postgres"
	"github.com/redis/go-redis/v9"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {

	config := config.InitConfig()

	pg, err := postgres.ConnectPostgres(&postgres.Config{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		DBName:   config.Database.Database,
		SSLMode:  "disable",
	})
	if err != nil {
		slog.Error("failed to connect to postgres", slog.String("error", err.Error()), slog.String("tag", "main"))
		panic(err)
	}

	slog.Info("dependencies initialized", slog.String("tag", "main"))

	db, err := database.InitDatabase(&database.Config{
		Dial: pg,
	})
	if err != nil {
		slog.Error("failed to initialize database", slog.String("error", err.Error()), slog.String("tag", "main"))
		panic(err)
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
	})
	err = redis.Ping(context.Background()).Err()
	if err != nil {
		slog.Error("failed to initialize redis", slog.String("error", err.Error()), slog.String("tag", "main"))
		panic(err)
	}

	router := routes.NewRouter(config, db, redis)
	router.RegisterRoutes()
}
