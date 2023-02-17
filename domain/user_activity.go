package domain

type UserActivity struct {
	User_ID    string `json:"user_id" bson:"user_id"`
	Action     string `json:"action" bson:"action"`
	Receiver   string `json:"action_on" bson:"action_on"`
	List_ID    string `json:"list_id" bson:"list_id"`
	Comment_ID string `json:"comment_id" bson:"comment_id"`
	Review_ID  string `json:"review_id" bson:"review_id"`
	Desc       string `json:"desc" bson:"desc"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  int64  `json:"updated_at" bson:"updated_at"`
	// what the user activities might look like

	// a user likes a list
	//a user puts a reaction on a list
	// a user likes their own list
	// a user put reaction on their own list
	// a user puts a review
	// a user puts a comment or like on someone else's review
	// for reaction we would need list id , as we will be viewing just the list and everything about it will
	// pop up .
	// what about userreviews though . ??? we would need to have that as an entity as well
	// should there be a user to book relationship first ?
	// named user books ?
	// then user-lists will just refer to the book object created by the user.
	// it will be a lot more convoluted.
	// So for now no reviews ?
	// or make reviews independent .
	// what about information about the list if a user wants to add ?
	//lets do  that atleast .
}

func NewUserActivity(user_id string, action string, receiver string, list_id string, comment_id string, review_id string, desc string, timestamp int64) *UserActivity {
	return &UserActivity{user_id, action, receiver, list_id, comment_id, review_id, desc, timestamp, timestamp}
}
