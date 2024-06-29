package main

import (
	"github.com/11SF/core-family-management/config"
	routes "github.com/11SF/core-family-management/pkg"
)

func main() {

	config := config.InitConfig()

	routes.NewRouter(config)
}
