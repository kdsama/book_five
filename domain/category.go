package domain

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Name              string
	Id                primitive.ObjectID
	SubCategories     []*Category
	Created_Timestamp int64
	Updated_Timestamp int64
}

func NewCategory(name string, id uuid.UUID, subcategory []*Category, created_timestamp int64, updated_timestamp int64) *Category {
	return &Category{name, id, subcategory, created_timestamp, updated_timestamp}
}
