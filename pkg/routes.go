package routes

import (
	"fmt"

	"github.com/11SF/core-family-management/config"
	corefamily "github.com/11SF/core-family-management/pkg/v1/family/core"
	"github.com/11SF/core-family-management/pkg/v1/family/handler"
	"github.com/11SF/core-family-management/pkg/v1/family/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type routes struct {
	config *config.Config
	db     *gorm.DB
	redis  *redis.Client
}

func NewRouter(config *config.Config, db *gorm.DB, redis *redis.Client) *routes {
	return &routes{
		config: config,
		db:     db,
		redis:  redis,
	}
}

func (r *routes) RegisterRoutes() {
	app := fiber.New(fiber.Config{
		AppName: r.config.AppName,
	})

	app.Use(cors.New())
	app.Use(requestid.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: fmt.Sprintf("%s Metrics Page", r.config.AppName)}))

	api := app.Group("/api")
	context := api.Group("/core-family-management")
	v1 := context.Group("/v1")

	dbFamily := repository.NewFamilyDB(r.db)
	redisFamily := repository.NewFamilyRedis(r.redis, r.config.RedisKeys)
	coreFamily := corefamily.NewService(dbFamily, redisFamily)
	createFamilyHandler := handler.NewCreateFamilyHandler(coreFamily.CreateFamily)
	getFamilyInfoHandler := handler.NewGetFamilyInfoHandler(coreFamily.GetFamilyInfo)
	getFamilyListHandler := handler.NewGetFamilyListHandler(coreFamily.GetFamilyList)

	v1.Post("/family", createFamilyHandler.Handler)
	v1.Get("/family/:familyId", getFamilyInfoHandler.Handler)
	v1.Get("/family", getFamilyListHandler.Handler)

	app.Listen(fmt.Sprintf(":%s", r.config.Server.Port))
}
