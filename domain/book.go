package domain

import (
	"github.com/kdsama/book_five/entity"
)

type Book struct {
	ID                string             `bson:"uuid"`
	Name              string             `bson:"name" json:"name"`
	Authors           []string           `bson:"authors" json:"authors"`
	Co_Authors        []string           `bson:"co_authors" json:"co_authors"`
	AudiobookUrls     []entity.UrlObject `bson:"audiobook_urls" json:"audiobook_urls"`
	EbookUrls         []entity.UrlObject `bson:"ebook_urls" json:"ebook_urls"`
	Hardcopies        []entity.UrlObject `bson:"hard_copies" json:"hard_copies"`
	Categories        []string           `bson:"categories" json:"categories"`
	Created_Timestamp int64              `bson:"createdAt" json:"createdAt"`
	Updated_Timestamp int64              `bson:"updatedAt" json:"updatedAt"`
	Verified          bool               `bson:"verified" json:"verified"`
}

func NewBook(name string, authors []string, co_authors []string, audio []entity.UrlObject, ebook []entity.UrlObject, hardcopy []entity.UrlObject, categories []string, timestamp int64) *Book {

	return &Book{Name: name, Authors: authors, Co_Authors: co_authors, AudiobookUrls: audio, EbookUrls: ebook, Hardcopies: hardcopy, Categories: categories, Created_Timestamp: timestamp}
}
