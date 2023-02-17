package repository

import "github.com/kdsama/book_five/domain"

type UserRepo interface {
	SaveUser(*domain.User) error
	GetUserByEmail(string) (*domain.User, error)
	GetUserByID(string) (*domain.User, error)
	CountUsersFromIDs([]string) (int64, error)
	GetUserNamesByIDs([]string) ([]string, error)
}

type UserRepository struct {
	UserRepo
}

func NewUserRepository(br UserRepo) *UserRepository {
	return &UserRepository{br}
}
