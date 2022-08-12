package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sport-test/entity"
)

func CreateNews(c *gin.Context, news entity.News) (string, error) {
	newsCollection := c.MustGet("newsDB").(*mongo.Collection)
	result, err := newsCollection.InsertOne(c, news)
	if err != nil {
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}
func GetNews(c *gin.Context) ([]entity.News, error) {
	var oneNews entity.News
	var manyNews []entity.News
	var cursor *mongo.Cursor
	var err error

	newsCollection := c.MustGet("newsDB").(*mongo.Collection)

	cursor, err = newsCollection.Find(c, bson.D{})
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

func GetOneNews(c *gin.Context, keyMongo string, valueMongo interface{}) (entity.News, error) {
	var news entity.News

	newsCollection := c.MustGet("newsDB").(*mongo.Collection)
	filter := bson.D{{keyMongo, valueMongo}}
	err := newsCollection.
		FindOne(context.TODO(), filter).
		Decode(&news)
	if err != nil {
		return news, err
	}
	return news, nil
}
