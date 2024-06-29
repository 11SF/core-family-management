package routes

import (
	"fmt"

	"github.com/11SF/core-family-management/config"
	coregoods "github.com/11SF/core-family-management/pkg/v1/goods/core"
	"github.com/11SF/core-family-management/pkg/v1/goods/handler"
	"github.com/11SF/core-family-management/pkg/v1/goods/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type routes struct {
	config *config.Config
}

func NewRouter(config *config.Config) *routes {
	return &routes{
		config: config,
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
	context := api.Group("/service-name")
	v1 := context.Group("/v1")

	dbGoods := repository.NewGoodsDB(nil)
	coreGoods := coregoods.NewService(dbGoods)
	getGoodsById := handler.NewGetGoodsListHandler(coreGoods.GetGoodsByID)

	v1.Get("/goods", getGoodsById.Handler)

	app.Listen(fmt.Sprintf(":%s", r.config.Server.Port))
}
