package service

import (
	"errors"
	"time"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

var (
	Err_Invalid_Hash = errors.New("hash provided is invalid")
)

func NewUserService(User repository.UserRepository) *UserService {

	return &UserService{User}
}

func (us *UserService) SaveUser(user string, password string) error {

	timestamp := time.Now().Unix()
	userObject := domain.NewUser(user, password, timestamp)
	us.UserRepo.SaveUser(userObject)
	return nil
}
