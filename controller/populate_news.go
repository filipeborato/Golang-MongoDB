package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"sport-test/entity"
	"sport-test/service"
	"strconv"
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
			newsArticleInformation, err := service.RequestNewsDetails(strconv.Itoa(newsItem.NewsArticleID))
			newsArticle := newsArticleInformation.News
			if err != nil {
				continue
			}
			published, _ := time.Parse("2006-01-02 15:04:05", newsItem.PublishDate)
			lastUpdated, _ := time.Parse("2006-01-02 15:04:05", newsArticle.LastUpdateDate)
			newsInsert := entity.News{
				NewsArticleID: newsItem.NewsArticleID,
				OptaMatchId:   &newsItem.OptaMatchId,
				Title:         newsItem.Title,
				Content:       newsArticle.BodyText,
				Teaser:        &newsItem.TeaserText,
				Url:           newsItem.ArticleURL,
				ImageUrl:      &newsArticle.ThumbnailImageURL,
				GalleryUrls:   &newsArticle.GalleryImageURLs,
				VideoUrl:      &newsArticle.VideoURL,
				Published:     published,
				LastUpdated:   lastUpdated,
			}
			id, err := service.CreateNews(c, newsInsert)
			if err != nil {
				continue
			}
			log.Print(id)

		}
	}
}
