package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"sport-test/entity"
	"sport-test/service"
)

func NewsMany(c *gin.Context) {
	//newsId, err := getNews(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Ticket não encontrado.", "erro:": err})
	//	return
	//}
	news, err := service.GetManyNews(c)

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

func TestGetNews(c *gin.Context) ([]entity.News, error) {
	var oneNews entity.News
	var manyNews []entity.News

	test := c.MustGet("test").(*mongo.Collection)
	cursor, err := test.Find(c, bson.D{})

	if err != nil {
		defer cursor.Close(c)
		return manyNews, err
	}

	for cursor.Next(c) {
		err := cursor.Decode(&oneNews)
		if err != nil {
			return manyNews, err
		}
		manyNews = append(manyNews, oneNews)
	}

	return manyNews, nil
}
