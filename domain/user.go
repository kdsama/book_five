package domain

import "github.com/kdsama/book_five/utils"

type User struct {
	ID        string `bson:"uuid" json:"user_id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"wwl" bson:"pwd"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

func NewUser(email string, name string, wwp string, timestamp int64) *User {

	return &User{ID: utils.GenerateUUID(), Email: email, Name: name, Password: wwp, CreatedAt: timestamp, UpdatedAt: timestamp}
}
