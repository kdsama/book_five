package domain

import (
	"github.com/kdsama/book_five/entity"
	"github.com/kdsama/book_five/utils"
)

type ListComment struct {
	ID        string          `json:"comment_id" bson:"uuid"`
	List_ID   string          `json:"list_id" bson:"list_id"`
	User_ID   string          `json:"user_id" bson:"user_id"` // Id of the user who made the comment
	Comment   string          `json:"comment" bson:"comment"`
	Reaction  entity.Reaction `json:"reaction" bson:"reaction"`
	CreatedAt int64           `json:"created_at" bson:"created_at"`
	UpdatedAt int64           `json:"updated_at" bson:"updated_at"`
}

func NewListComment(list_id string, user_id string, comment string, timestamp int64) *ListComment {

	return &ListComment{utils.GenerateUUID(), list_id, user_id, comment, *entity.NewReaction(), timestamp, timestamp}
}
