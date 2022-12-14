package main

import (
	"github.com/gin-gonic/gin"
	"sport-test/controller"
	"sport-test/middleware"
	"sport-test/schedule"
)

func main() {
	if !schedule.CronjobCommands() {
		server()
	}
}

func server() {
	version := "v1"
	server := gin.Default()
	api := server.Group(version+"/api", middleware.Database())
	{
		api.GET("/news", controller.NewsMany)
		api.GET("/news/:news_id", controller.NewsOne)
		api.GET("/populate-test", controller.PopulateNewsInformation)
	}
	//api.GET("/one", func(ctx *gin.Context) {
	//	ctx.JSON(200, controller.PopulateInformation)
	//})
	server.Run(":8080")
}