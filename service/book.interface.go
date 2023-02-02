package service

type BookServiceInterface interface {
	SaveBook(name string, authors []string, co_authors []string, audiobook_urls []string, ebook_urls []string, hard_copies []string, categories []string) error
}

type BookDI struct {
	BookServiceInterface
}

func NewBookServiceInterface(br BookServiceInterface) *BookDI {
	return &BookDI{br}
}
