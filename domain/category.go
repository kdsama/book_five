package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	Name              string
	Id                uuid.UUID
	SubCategories     []*Category
	Created_Timestamp int64
	Updated_Timestamp int64
}

var categories = []*Category{}
var err_CategoryNotFouund = errors.New("Category is not present ")

func NewCategory(name string, id uuid.UUID, subcategory []*Category, created_timestamp int64, updated_timestamp int64) *Category {
	return &Category{name, id, subcategory, created_timestamp, updated_timestamp}
}
func (c *Category) saveCategory() bool {
	categories = append(categories, c)
	return true

}
func (c *Category) getCategories() []*Category {
	return categories
}

func (c *Category) getCategoryByUUID(id uuid.UUID) (*Category, error) {
	for _, category := range categories {
		if category.Id == id {
			return id, nil
		}
	}
	return &Category{}, err_CategoryNotFouund
}
