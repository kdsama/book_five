package repository

import "github.com/kdsama/book_five/domain"

type UserRepo interface {
	SaveUser(*domain.User) error
	GetUserByEmail(string) (*domain.User, error)
	GetUserIdByName(string) (string, error)
	GetUserById(string) (*domain.User, error)
}

type UserRepository struct {
	UserRepo
}

func NewUserRepository(br UserRepo) *UserRepository {
	return &UserRepository{br}
}
