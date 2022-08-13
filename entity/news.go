package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type News struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	NewsArticleID int                `bson:"newsArticleId"`
	TeamId        string             `bson:"teamId"`
	OptaMatchId   *string            `bson:"optaMatchId"`
	Title         string             `bson:"title"`
	Teaser        *string            `bson:"teaser"`
	Content       string             `bson:"content"`
	Url           string             `bson:"url"`
	ImageUrl      *string            `bson:"imageUrl"`
	GalleryUrls   *string            `bson:"galleryUrls"`
	VideoUrl      *string            `bson:"videoUrl"`
	Published     time.Time          `bson:"published"`
	LastUpdated   time.Time          `bson:"last_updated"`
}

type Metadata struct {
	CreatedAt  time.Time `json:"createdAt"`
	TotalItems int       `json:"totalItems"`
	Sort       string    `json:"sort"`
}
