package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roundrobinquantum/account-api/config"
	"github.com/roundrobinquantum/account-api/handler"
	 clt "github.com/roundrobinquantum/api-client/client"
)

func main() {
	cfg := config.GetConfigFromEnvVariable()

	engine := gin.New()
	engine.Use(gin.Recovery())

	restClient:= clt.NewClient(cfg.ConnTimeout,cfg.ReadTimeout,cfg.WriteTimeout)

	handler.InitializeRoutes(engine,restClient,cfg.Uri)

	engine.Run(":8080")
}
