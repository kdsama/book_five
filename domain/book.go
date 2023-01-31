package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID   `bson:"_id"`
	Name              string               `bson:"name" json:"name"`
	Authors           []string             `bson:"authors" json:"authors"`
	Co_Authors        []string             `bson:"co_authors" json:"co_authors"`
	AudiobookUrls     []string             `bson:"audiobook_urls" json:"audiobook_urls"`
	EbookUrls         []string             `bson:"ebook_urls" json:"ebook_urls"`
	Hardcopies        []string             `bson:"hard_copies" json:"hard_copies"`
	Categories        []primitive.ObjectID `bson:"categories" json:"categories"`
	Created_Timestamp int64                `bson:"createdAt" json:"createdAt"`
	Updated_Timestamp int64                `bson:"updatedAt" json:"updatedAt"`
}

var book_Collection = []Book{}

func NewBook(name string, authors []string, co_authors []string, audio []string, ebook []string, hardcopy []string, categories []primitive.ObjectID, timestamp int64) *Book {

	return &Book{Name: name, Authors: authors, Co_Authors: co_authors, AudiobookUrls: audio, EbookUrls: ebook, Hardcopies: hardcopy, Categories: categories, Created_Timestamp: timestamp}
}
