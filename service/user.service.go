package service

import (
	"errors"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

var MIN_PASSWORD_HASH_LENGTH int = 30
var (
	Err_Invalid_Hash = errors.New("hash provided is invalid")
)

func NewUserService(User repository.UserRepository) *UserService {

	return &UserService{User}
}

func (us *UserService) SaveUser(user *domain.User) error {
	if len(user.Password) < MIN_PASSWORD_HASH_LENGTH {
		return Err_Invalid_Hash
	}
	return nil
}
