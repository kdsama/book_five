package domain

import (
	"github.com/kdsama/book_five/entity"
)

type ListComment struct {
	List_ID  string          `json:"list_id" bson:"list_id"`
	User_ID  string          `json:"user_id" bson:"user_id"` // Id of the user who made the comment
	Comment  string          `json:"comment" bson:"comment"`
	Reaction entity.Reaction `json:"reaction" bson:"reaction"`
}
