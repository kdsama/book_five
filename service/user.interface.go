package service

import domain "github.com/kdsama/book_five/domain"

type UserServiceInterface interface {
	SaveUser(string, string, string) (string, error)
	LoginUser(string, string) (string, error)
	CountUsersFromIDs([]string) (int64, error)
	GetUserByID(string) (*domain.User, error)
	GetUserNamesByIDs([]string) ([]string, error)
}

type UserDI struct {
	UserServiceInterface
}

func NewUserServiceInterface(br UserServiceInterface) *UserDI {
	return &UserDI{br}
}
