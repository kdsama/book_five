package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryServiceInterface interface {
	SaveCategory(name string, categories []string) error
	GetIdsByNames(names []string) ([]primitive.ObjectID, error)
}

type CategoryDI struct {
	CategoryServiceInterface
}

func NewCategoryServiceInterface(br CategoryServiceInterface) *CategoryDI {
	return &CategoryDI{br}
}
