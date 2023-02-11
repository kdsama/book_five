package service

import domain "github.com/kdsama/book_five/domain"

type BookServiceInterface interface {
	SaveBook(name string, authors []string, co_authors []string, audiobook_urls []string, ebook_urls []string, hard_copies []string, categories []string) error
	UpsertBooksAndGetIDs([]domain.Book) ([]string, []error)
}

type BookDI struct {
	BookServiceInterface
}

func NewBookServiceInterface(br BookServiceInterface) *BookDI {
	return &BookDI{br}
}
