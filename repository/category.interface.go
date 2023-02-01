package repository

import (
	"github.com/kdsama/book_five/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepo interface {
	SaveCategory(*domain.Category) error
	GetCategoryByID(primitive.ObjectID) (*domain.Category, error)
	GetCategoriesByManyIDs([]primitive.ObjectID) ([]domain.Category, error)
	GetCategoriesByNames([]string) ([]domain.Category, error)
	GetIdsByNames([]string) ([]primitive.ObjectID, error)
}

type CategoryRepository struct {
	CategoryRepo
}

func NewCategoryRepository(br CategoryRepo) *CategoryRepository {
	return &CategoryRepository{br}
}
