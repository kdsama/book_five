package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id                primitive.ObjectID   `bson:"_id"`
	Name              string               `bson:"name" json:"name"`
	SubCategories     []primitive.ObjectID `bson:"sub_categories" json:"sub_categories"`
	Created_Timestamp int64                `bson:"created_timestamp" json:"created_timestamp"`
	Updated_Timestamp int64                `bson:"updated_timestamp" json:"updated_timestamp"`
}

func NewCategory(name string, subcategory []primitive.ObjectID, timestamp int64) *Category {
	return &Category{Name: name,
		SubCategories:     subcategory,
		Created_Timestamp: timestamp,
		Updated_Timestamp: timestamp,
	}
}
