package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"sport-test/entity"
	"sport-test/service"
	"time"
)

func PopulateNewsInformation(c *gin.Context) {
	newsInformation, err := service.RequestNewsInformation()
	if err != nil {
		log.Print(err)
		return
	}
	newsLetter := newsInformation.NewsletterNewsItems.NewsletterNewsItem

	for _, newsItem := range newsLetter {

		news, _ := service.GetOneNews(c, "newsArticleId", newsItem.NewsArticleID)

		if (entity.News{}) == news {
			published, _ := time.Parse("2006-01-02", newsItem.PublishDate)
			newsInsert := entity.News{
				NewsArticleID: newsItem.NewsArticleID,
				OptaMatchId:   &newsItem.OptaMatchId,
				Title:         newsItem.Title,
				Teaser:        &newsItem.TeaserText,
				Url:           newsItem.ArticleURL,
				Published:     published,
			}
			result, err := service.CreateNews(c, newsInsert)
			if err != nil {
				continue
			}
			log.Print(result)

		}
	}
}
