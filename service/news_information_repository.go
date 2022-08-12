package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sport-test/entity"
)

func GetNews(c *gin.Context, keyMongo string, valueMongo interface{}) ([]entity.News, error) {
	var oneNews entity.News
	var manyNews []entity.News
	var cursor *mongo.Cursor

	newsCollection := c.MustGet("newsDB").(*mongo.Collection)
	if keyMongo == "" && valueMongo == nil {
		cursor, err := newsCollection.Find(c, bson.D{})
		if err != nil {
			defer cursor.Close(c)
			return manyNews, err
		}

	} else {
		cursor, err := newsCollection.Find(c, bson.D{{keyMongo, valueMongo}})
		if err != nil {
			defer cursor.Close(c)
			return manyNews, err
		}
	}
	if cursor == nil {
		return manyNews, nil
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

func GetOneNewsInformation(c *gin.Context, id string) (entity.News, error) {
	var news entity.News
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return news, err
	}
	newsCollection := c.MustGet("newsDB").(*mongo.Collection)

	err = newsCollection.
		FindOne(c, bson.D{{"_id", objectId}}).
		Decode(&news)
	if err != nil {
		return news, err
	}
	return news, nil
}
