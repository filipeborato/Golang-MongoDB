package entity

import "encoding/xml"

type NewsArticleInformation struct {
	XMLName xml.Name    `xml:"NewsArticleInformation"`
	News    NewsArticle `xml:"NewsArticle"`
}

type NewsArticle struct {
	XMLName           xml.Name `xml:"NewsArticle"`
	ArticleURL        string   `xml:"ArticleURL"`
	NewsArticleID     int      `xml:"NewsArticleID"`
	PublishDate       string   `xml:"PublishDate"`
	Taxonomies        string   `xml:"Taxonomies"`
	TeaserText        string   `xml:"TeaserText"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
	BodyText          string   `xml:"BodyText"`
	GalleryImageURLs  string   `xml:"GalleryImageURLs"`
	VideoURL          string   `xml:"VideoURL"`
	Title             string   `xml:"Title"`
	OptaMatchId       string   `xml:"OptaMatchId"`
	LastUpdateDate    string   `xml:"LastUpdateDate"`
	IsPublished       bool     `xml:"IsPublished"`
}
