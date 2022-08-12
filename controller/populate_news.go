package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sport-test/service"
)

func PopulateNewsInformation(c *gin.Context) {
	newsInformation, err := service.RequestNewsInformation()
	if err != nil {
		log.Print(err)
		return
	}
	newsLetter := newsInformation.NewsletterNewsItems.NewsletterNewsItem

	for _, newsItem := range newsLetter {
		fmt.Printf("newsItem %+v\n", newsItem)
		news, err := service.GetNews(c, "", nil)
		if err != nil {
			return
		}
		if news == nil {
			log.Print("new nil", news)
			continue
		}
	}
}
