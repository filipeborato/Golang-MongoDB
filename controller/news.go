package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sport-test/service"
	"strconv"
)

func NewsMany(c *gin.Context) {
	news, err := service.GetNews(c)
	if err != nil {
		log.Print(err)
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}
func NewsOne(c *gin.Context) {
	newsId := c.Param("news_id")
	newsIdInt, _ := strconv.Atoi(newsId)
	news, err := service.GetOneNews(c, "newsArticleId", newsIdInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "the News was not find.", "erro:": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": news})

}
