package entity

import "encoding/xml"

type NewListInformation struct {
	XMLName             xml.Name            `xml:"NewListInformation"`
	NewsletterNewsItems NewsletterNewsItems `xml:"NewsletterNewsItems"`
}

type NewsletterNewsItems struct {
	XMLName            xml.Name             `xml:"NewsletterNewsItems"`
	NewsletterNewsItem []NewsletterNewsItem `xml:"NewsletterNewsItem"`
}
type NewsletterNewsItem struct {
	XMLName           xml.Name `xml:"NewsletterNewsItem"`
	ArticleURL        string   `xml:"ArticleURL"`
	NewsArticleID     int      `xml:"NewsArticleID"`
	PublishDate       string   `xml:"PublishDate"`
	Taxonomies        string   `xml:"Taxonomies"`
	TeaserText        string   `xml:"TeaserText"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
	Title             string   `xml:"Title"`
	OptaMatchId       string   `xml:"OptaMatchId"`
	LastUpdateDate    string   `xml:"LastUpdateDate"`
	IsPublished       bool     `xml:"IsPublished"`
}
