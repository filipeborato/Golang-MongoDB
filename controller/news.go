package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"sport-test/service"
)

func NewsMany(c *gin.Context) {
	news, err := service.GetNews(c)
	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fault",
			"message": "the News was not find.",
			"erro:":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"news":   news,
	})
}
func NewsOne(c *gin.Context) {
	newsId := c.Param("news_id")
	newsIdObject, _ := primitive.ObjectIDFromHex(newsId)
	news, err := service.GetOneNews(c, "_id", newsIdObject)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fault",
			"message": "the News was not find.",
			"erro:":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"news":   news,
	})

}
