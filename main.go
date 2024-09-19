package main

import (
	"github.com/gin-gonic/gin"
	"url-shortener--go-gin/common/app"
	"url-shortener--go-gin/common/postgresql"
	"url-shortener--go-gin/controller"
	"url-shortener--go-gin/persistence"
	"url-shortener--go-gin/service"
)

func main() {
	server := gin.Default()

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
