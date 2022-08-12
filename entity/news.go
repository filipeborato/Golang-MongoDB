package entity

import "time"

type News struct {
	NewsArticleID int       `bson:"newsArticleId"`
	TeamId        string    `bson:"teamId"`
	OptaMatchId   *string   `bson:"optaMatchId"`
	Title         string    `bson:"title"`
	Teaser        *string   `bson:"teaser"`
	Content       string    `bson:"content"`
	Url           string    `bson:"url"`
	ImageUrl      *string   `bson:"imageUrl"`
	GalleryUrls   *string   `bson:"galleryUrls"`
	VideoUrl      *string   `bson:"videoUrl"`
	Published     time.Time `json:"published" bson:"published"`
	LastUpdated   time.Time `json:"last_updated" bson:"last_updated"`
}
