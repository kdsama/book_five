package service

import (
	"errors"
	"fmt"
	"time"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/entity"
	"github.com/kdsama/book_five/repository"
)

var (
	Err_Invalid_Categories = errors.New("categories are invalid")
	Err_Invalid_Category   = errors.New("category is invalid")
)

type BookService struct {
	bookRepo   repository.BookRepository
	catService CategoryDI
}

func NewBookService(book repository.BookRepository, catService CategoryDI) *BookService {
	return &BookService{book, catService}
}
func (bs *BookService) SaveBook(name string, image_url string, authors []string, co_authors []string, audiobook_urls []string, ebook_urls []string, hard_copies []string, categories []string) error {
	// validation checks already done in the http handler .
	// We just need to save it for now.
	timestamp := time.Now().Unix()
	// Get all Ids of categories
	categoryIds, err := bs.catService.GetIdsByNames(categories)
	if err != nil {
		return err
	}
	if len(categoryIds) == 0 {
		fmt.Println(categoryIds)
		if len(categories) == 1 {
			return Err_Invalid_Category
		}
		return Err_Invalid_Categories
	}
	AudioBookObjects := []entity.UrlObject{}
	EbookObjects := []entity.UrlObject{}
	HardCopyObjects := []entity.UrlObject{}
	for i := range audiobook_urls {
		AudioBookObjects = append(AudioBookObjects, *entity.MakeUrlObject(audiobook_urls[i]))
	}
	for j := range ebook_urls {
		EbookObjects = append(EbookObjects, *entity.MakeUrlObject(ebook_urls[j]))
	}
	for k := range hard_copies {
		HardCopyObjects = append(HardCopyObjects, *entity.MakeUrlObject(hard_copies[k]))
	}

	BookObject := domain.NewBook(name, image_url, authors, co_authors, AudioBookObjects, EbookObjects, HardCopyObjects, categoryIds, timestamp)
	err = bs.bookRepo.SaveBook(BookObject)
	if err != nil {
		return err
	}
	return nil
}
