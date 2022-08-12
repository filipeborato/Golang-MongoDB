package test

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sport-test/controller"
	"sport-test/middleware"
	"sport-test/service"
	"testing"
)

func TestPopulate(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	client := middleware.MongoDB()
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
	controller.PopulateNewsInformation(c)

}
func TestNewsDetails(test *testing.T) {
	service.RequestNewsDetails("418271")
}
func TestNewsInformation(test *testing.T) {
	service.RequestNewsInformation()
}
