package service

import (
	"fmt"
	"time"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/entity"
	"github.com/kdsama/book_five/repository"
)

type BookService struct {
	bookRepo   repository.BookRepository
	catService CategoryService
}

func NewBookService(book repository.BookRepository, catService CategoryService) *BookService {
	return &BookService{book, catService}
}
func (bs *BookService) SaveBook(name string, authors []string, co_authors []string, audiobook_urls []string, ebook_urls []string, hard_copies []string, categories []string) error {
	// validation checks already done in the http handler .
	// We just need to save it for now.
	timestamp := time.Now().Unix()
	// Get all Ids of categories
	categoryIds, err := bs.catService.GetIdsByNames(categories)
	if err != nil {
		return err
	}
	AudioBookObjects := []entity.UrlObject{}
	EbookObjects := []entity.UrlObject{}
	HardCopyObjects := []entity.UrlObject{}
	for i := range audiobook_urls {
		fmt.Println(audiobook_urls)
		AudioBookObjects = append(AudioBookObjects, *entity.MakeUrlObject(audiobook_urls[i]))
	}
	for j := range ebook_urls {
		EbookObjects = append(EbookObjects, *entity.MakeUrlObject(ebook_urls[j]))
	}
	for k := range hard_copies {
		HardCopyObjects = append(HardCopyObjects, *entity.MakeUrlObject(hard_copies[k]))
	}
	fmt.Println(AudioBookObjects)
	BookObject := domain.NewBook(name, authors, co_authors, AudioBookObjects, EbookObjects, HardCopyObjects, categoryIds, timestamp)
	err = bs.bookRepo.Client.SaveBook(BookObject)
	if err != nil {
		return err
	}
	return nil
}

// func (bs *BookService) GetBookByName(name string) (*domain.Book, error) {

// }
