package controller

import (
	"github.com/gin-gonic/gin"
	"sport-test/service"
)

func PopulateInformation(*gin.Context) {
	service.RequestNewsInformation()
}

func NewsMany(c *gin.Context) {
	//newsId, err := getNews(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Ticket não encontrado.", "erro:": err})
	//	return
	//}
}
func NewsOne(c *gin.Context) {
	//newsId, err := getNews(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Ticket não encontrado.", "erro:": err})
	//	return
	//}
}
