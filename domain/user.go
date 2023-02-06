package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"wwl" bson:"pwd"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}

func NewUser(email string, wwp string, timestamp int64) *User {

	return &User{Email: email, Password: wwp, CreatedAt: timestamp, UpdatedAt: timestamp}
}
