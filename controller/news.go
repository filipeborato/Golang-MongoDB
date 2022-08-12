package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sport-test/service"
)

func NewsMany(c *gin.Context) {
	//newsId, err := getNews(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Ticket não encontrado.", "erro:": err})
	//	return
	//}
	news, err := service.GetNews(c, "", nil)

	if err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{"news": news})
}
func NewsOne(c *gin.Context) {
	//newsId, err := getNews(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Ticket não encontrado.", "erro:": err})
	//	return
	//}
}
