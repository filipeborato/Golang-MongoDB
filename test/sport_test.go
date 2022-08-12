package test

import (
	"github.com/gin-gonic/gin"
	"sport-test/controller"
	"sport-test/middleware"
	"testing"
)

func TestSport(test *testing.T) {
	var c *gin.Context
	middleware.Database()
	controller.PopulateNewsInformation(c)
}
