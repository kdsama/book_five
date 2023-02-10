package domain

type UserActivity struct {
	User_ID    string `json:"user_id" bson:"user_id"`
	Action     string `json:"action" bson:"action"`
	Receiver   string `json:"action_on" bson:"action_on"`
	List_ID    string `json:"list_id" bson:"list_id"`
	Comment_ID string `json:"comment_id" bson:"comment_id"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  int64  `json:"updated_at" bson:"updated_at"`
	// will have id of the user
	// last time user added a book. We would look up to latest entry to check whether a user is allowed to add another book or not .
	// should probably have the activities the user do on other people-lists .
	// should have timestamp of the particular activity
	// probably will be different documents
	//
}

func NewUserActivity(user_id string, action string, receiver string, list_id string, comment_id string, timestamp int64) *UserActivity {
	return &UserActivity{user_id, action, receiver, list_id, comment_id, timestamp, timestamp}
}
