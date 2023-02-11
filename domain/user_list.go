package domain

import (
	"github.com/kdsama/book_five/entity"
)

type UserList struct {
	// what a user would like to put on as lists
	// there will be different lists
	// list can be distinguished by categories
	// list can tell which author was there the most
	// list consists of books/ids
	// lists have a type
	// one can add a custom name to the list
	// list can be liked or disliked by others ?
	// might still need the books the user didnot like
	// multiple lists can have the same books ?
	// someone can comment on a list
	// someone can X,Y,Z a list, according to the options provided.
	// people cannot rate a list
	// items cannot be added to a list more than 1 time a day
	// Items cannot be added to any list for a particular user more than 1 time a day. This will be checked by checking the updated at of the particular list.
	// And then Check in User activity . Also means whenever we update updatedAt , we would also need to update UserActivity
	// we should be able to figure out if a person have read any book recently
	// people can subscribe to a list of a particular user ?
	// its controlled on how many lists a particular user can make according to some calculations
	// for now a blank function which returns 5 , maybe
	// Length of the list cannot be more than 5 as well

	ID        string          `bson:"uuid"`
	User_ID   string          `bson:"user_id"`
	Book_IDs  []string        `json:"book_ids" bson:"book_ids"`
	Name      string          `json:"name" bson:"name"`
	Reactions entity.Reaction `json:"reaction" bson:"reaction"`
	Comments  []ListComment   `json:"comments" bson:"comments"`
	CreatedAt int64           `json:"created_at" bson:"created_at"`
	UpdatedAt int64           `json:"updated_at" bson:"updated_at"`
}

func NewUserList(user_id string, book_ids []string, name string, timestamp int64) *UserList {

	return &UserList{"", user_id, book_ids, name, *entity.NewReaction(), []ListComment{}, timestamp, timestamp}
}
