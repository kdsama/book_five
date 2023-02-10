package domain

type Category struct {
	Id                string   `bson:"_id"`
	Name              string   `bson:"name" json:"name"`
	SubCategories     []string `bson:"sub_categories" json:"sub_categories"`
	Created_Timestamp int64    `bson:"created_timestamp" json:"created_timestamp"`
	Updated_Timestamp int64    `bson:"updated_timestamp" json:"updated_timestamp"`
}

func NewCategory(name string, subcategory []string, timestamp int64) *Category {
	return &Category{Name: name,
		SubCategories:     subcategory,
		Created_Timestamp: timestamp,
		Updated_Timestamp: timestamp,
	}
}
