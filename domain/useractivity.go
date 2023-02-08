package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserActivity struct {
	User_ID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Action     entity.Action      `json:"action" bson:"action"`
	ActionOn   entity.ActionOn    `json:"action_on" bson:"action_on"`
	Subject_ID entity.Subject_ID  `json:"subject_id" bson:"subject_id"`

	// will have id of the user
	// last time user added a book. We would look up to latest entry to check whether a user is allowed to add another book or not .
	// should probably have the activities the user do on other people-lists .
	// should have timestamp of the particular activity
	// probably will be different documents
	//
}
