package repository

import (
	"github.com/kdsama/book_five/domain"
)

type CategoryRepo interface {
	SaveCategory(*domain.Category) error
	GetCategoryByID(string) (*domain.Category, error)
	GetCategoriesByManyIDs([]string) ([]domain.Category, error)
	GetCategoriesByNames([]string) ([]domain.Category, error)
	GetIdsByNames([]string) ([]string, error)
}

type CategoryRepository struct {
	CategoryRepo
}

func NewCategoryRepository(br CategoryRepo) *CategoryRepository {
	return &CategoryRepository{br}
}
