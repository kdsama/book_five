package repository

import (
	"github.com/kdsama/book_five/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepo interface {
	SaveCategory(*domain.Category) error
	GetCategoryByID(primitive.ObjectID) *domain.Category
	GetCategoriesByManyIDs([]primitive.ObjectID) []*domain.Category
}

type CategoryRepository struct {
	repo CategoryRepo
}

func NewCategoryRepository(br CategoryRepo) *CategoryRepository {
	return &CategoryRepository{br}
}
