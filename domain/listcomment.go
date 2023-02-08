package domain

import (
	"github.com/kdsama/book_five/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListComment struct {
	List_ID  primitive.ObjectID `json:"list_id" bson:"list_id"`
	User_ID  primitive.ObjectID `json:"user_id" bson:"user_id"` // Id of the user who made the comment
	Comment  string             `json:"comment" bson:"comment"`
	Reaction entity.Reaction    `json:"reaction" bson:"reaction"`
}
