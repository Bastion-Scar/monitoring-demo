package main

import (
	"monitoring-demo/internal/logger"
	"monitoring-demo/internal/middlewares"
	"monitoring-demo/internal/service"
	"monitoring-demo/internal/storage"
	"monitoring-demo/models"

	"github.com/gin-gonic/gin"
)

//TODO: ДОПОЛНИТЕЛЬНО написать Dockerfile и Docker-compose

func main() {
	zapLogger := logger.Logger()
	initStorage := storage.NewMySQLStorage()
	userService := service.NewUserService(initStorage)

	r := gin.New()
	db := storage.InitDB()

	r.Use(middlewares.LoggerMiddleware(zapLogger))
	r.Use(gin.Recovery())
	r.POST("/shorten", func(c *gin.Context) {
		userService.Save(c)
	})

	r.GET("/:code", func(c *gin.Context) {
		userService.Load(c)

	})

	go models.SendLogs(db)

	if err := r.Run(":8080"); err != nil {
		zapLogger.Fatal("Ошибка при запуске сервера")
	}

}
