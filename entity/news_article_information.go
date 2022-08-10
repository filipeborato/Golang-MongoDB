package entity

import "encoding/xml"

type NewsArticleInformation struct {
	XMLName xml.Name    `xml:"NewsArticleInformation"`
	News    NewsArticle `xml:"NewsArticle"`
}

type NewsArticle struct {
	XMLName          xml.Name `xml:"NewsArticle"`
	NewsArticleID    int      `xml:"NewsArticleID"`
	BodyText         string   `xml:"BodyText"`
	GalleryImageURLs string   `xml:"GalleryImageURLs"`
	VideoURL         string   `xml:"VideoURL"`
}
