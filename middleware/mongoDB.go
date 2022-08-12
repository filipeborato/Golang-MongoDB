package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func Database() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := MongoDB()
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
		c.Set("mongoDB", client)
		db := client.Database("sport")
		c.Set("sportDB", db)
		c.Set("newsDB", db.Collection("news_information"))
		c.Set("test", db.Collection("test"))
		c.Next()
	}
}

func MongoDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Println("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Print(err)
	}
	return client
}
