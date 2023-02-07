package service

import (
	"errors"
	"time"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type UserService struct {
	UserRepo repository.UserRepository
}

var (
	Err_Invalid_Hash = errors.New("hash provided is invalid")
	Err_User_Present = errors.New("user already present")
)

func NewUserService(User repository.UserRepository) *UserService {

	return &UserService{User}
}

func (us *UserService) SaveUser(email string, password string) error {

	_, err := us.UserRepo.GetUserByEmail(email)

	if err == nil {
		return Err_User_Present
	}

	if err != repository.Err_UserNotFound {
		return err
	}

	timestamp := time.Now().Unix()
	encryptedPassword, err := utils.GenerateHashForPassword(password)
	if err != nil {
		return err
	}
	userObject := domain.NewUser(email, encryptedPassword, timestamp)
	us.UserRepo.SaveUser(userObject)
	return nil
}
