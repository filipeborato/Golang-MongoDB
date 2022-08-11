package entity

import "time"

type News struct {
	Id1         string      `bson:"id"`
	TeamId      string      `bson:"teamId"`
	OptaMatchId interface{} `bson:"optaMatchId"`
	Title       string      `bson:"title"`
	Type        []string    `bson:"type"`
	Teaser      interface{} `bson:"teaser"`
	Content     string      `bson:"content"`
	Url         string      `bson:"url"`
	ImageUrl    string      `bson:"imageUrl"`
	GalleryUrls interface{} `bson:"galleryUrls"`
	VideoUrl    interface{} `bson:"videoUrl"`
	Published   time.Time   `bson:"published"`
}
