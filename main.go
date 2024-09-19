package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"url-shortener--go-gin/common/app"
	"url-shortener--go-gin/common/postgresql"
	"url-shortener--go-gin/controller"
	"url-shortener--go-gin/controller/middlewares"
	"url-shortener--go-gin/persistence"
	"url-shortener--go-gin/service"
)

func main() {
	bucket := middlewares.NewTokenBucket(2, 1, time.Second*10)
	server := gin.Default()
	server.Use(middlewares.RateLimiter(bucket))

	configurationManager := app.NewConfigurationManager()
	db := postgresql.GetConnection(configurationManager.PostgreSqlConfig)

	postgresql.MigrateTables(db)

	urlRepository := persistence.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)

	urlController.RegisterUrlRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
