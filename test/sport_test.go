package test

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sport-test/controller"
	"sport-test/middleware"
	"sport-test/service"
	"testing"
)

func TestPopulate(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c = middleware.TestMongo(c)
	controller.PopulateNewsInformation(c)

}
func TestNewsDetails(test *testing.T) {
	service.RequestNewsDetails("418271")
}
func TestNewsInformation(test *testing.T) {
	service.RequestNewsInformation()
}

func TestNewsMany(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c = middleware.TestMongo(c)
	controller.NewsMany(c)
}

func TestNewsOne(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c = middleware.TestMongo(c)
	c.Set("news_id", "418271")

	controller.NewsMany(c)
}
