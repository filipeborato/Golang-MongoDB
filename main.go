package main

import (
	"github.com/gin-gonic/gin"
	"sport-test/controller"
	"sport-test/middleware"
)

func main() {
	version := "v1"
	server := gin.Default()
	mongoDB := middleware.Database()
	api := server.Group(version+"/api", mongoDB)
	api.GET("/news", controller.NewsMany)
	api.GET("/news/:news_id", controller.NewsOne)

	//api.GET("/one", func(ctx *gin.Context) {
	//	ctx.JSON(200, controller.PopulateInformation)
	//})
	server.Run(":8080")
}
