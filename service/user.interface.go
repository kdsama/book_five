package service

import domain "github.com/kdsama/book_five/domain"

type UserServicer interface {
	SaveUser(string, string, string) (string, error)
	LoginUser(string, string) (string, error)
	CountUsersFromIDs([]string) (int64, error)
	GetUserByID(string) (*domain.User, error)
	GetUserNamesByIDs([]string) ([]string, error)
}
