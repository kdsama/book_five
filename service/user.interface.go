package service

import "github.com/kdsama/book_five/domain"

type UserServiceInterface interface {
	SaveUser(string, string) error
	GetUserById(string) (*domain.User, error)
}

type UserDI struct {
	UserServiceInterface
}

func NewUserServiceInterface(br UserServiceInterface) *UserDI {
	return &UserDI{br}
}
