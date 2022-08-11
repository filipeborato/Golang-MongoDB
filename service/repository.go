package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sport-test/entity"
)

func GetManyNews(c *gin.Context) ([]entity.News, error) {
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
