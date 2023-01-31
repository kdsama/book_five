package services

import (
	"time"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type BookService struct {
	bookRepo        repository.BookRepository
	categoryService service.categoryService
}

func NewBookService(book repository.BookRepository, categoryRepo repository.CategoryRepository) *BookService {
	return &BookService{book, categoryRepo}
}
func (bs *BookService) SaveBook(name string, authors []string, co_authors []string, audio []string, ebook []string, onlinecopy []string, categories []string) {
	// validation checks already done in the http handler .
	// We just need to save it for now.
	timestamp := time.Now().Unix()

	BookObject := domain.NewBook(name, authors, co_authors, audio, ebook, onlinecopy, categories, timestamp)
	bs.bookRepo.SaveBook(BookObject)

}

// func (bs *BookService) GetBookByName(name string) (*domain.Book, error) {

// }
