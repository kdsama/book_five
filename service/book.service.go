package service

import (
	"time"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type BookService struct {
	bookRepo   repository.BookRepository
	catService CategoryService
}

func NewBookService(book repository.BookRepository, catService CategoryService) *BookService {
	return &BookService{book, catService}
}
func (bs *BookService) SaveBook(name string, authors []string, co_authors []string, audio []string, ebook []string, hardcopy []string, categories []string) error {
	// validation checks already done in the http handler .
	// We just need to save it for now.
	timestamp := time.Now().Unix()
	// Get all Ids of categories
	categoryIds, err := bs.catService.GetIdsByNames(categories)
	if err != nil {
		return err
	}
	BookObject := domain.NewBook(name, authors, co_authors, audio, ebook, hardcopy, categoryIds, timestamp)
	err = bs.bookRepo.Client.SaveBook(BookObject)
	if err != nil {
		return err
	}
	return nil
}

// func (bs *BookService) GetBookByName(name string) (*domain.Book, error) {

// }
