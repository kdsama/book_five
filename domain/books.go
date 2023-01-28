package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id 				  uuid.UUID
	Name              string
	Audio             []string 
	Ebook             []string 
	Onlinecopy        []string
	Category 		  []uuid.UUID 
	Created_Timestamp int64
	Updated_Timestamp int64
}

var Book_Collection = []Book{}
var err_BookNotFound = errors.New("Book couldnot be found ")

func NewBook(name string, audio []string, ebook []string, onlinecopy []string) *Book {
	timestamp := time.Now().Unix()
	return &Book{Name: name, Audio: audio, Ebook: ebook, Onlinecopy: onlinecopy, Created_Timestamp: timestamp}
}

func (b *Book) SaveBook() bool {
	Book_Collection = append(Book_Collection, *b)
	return true
}

func (b *Book) GetBookByName(name string) (*Book, error) {
	for _, book := range Book_Collection {
		if book.Name == name {
			return &book, nil
		}
	}
	return &Book{}, err_BookNotFound
}

func (b *Book) GetBooksBy