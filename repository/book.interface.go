package repository

import "github.com/kdsama/book_five/domain"

type BookRepo interface {
	SaveBook(*domain.Book) error
}

type BookRepository struct {
	repo BookRepo
}

func NewBookRepository(br BookRepo) *BookRepository {
	return &BookRepository{br}
}
