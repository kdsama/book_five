package service

import (
	"errors"
	"time"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type UserService struct {
	UserRepo         repository.UserRepo
	UserTokenService UserTokenInterface
}

var (
	Err_Invalid_Hash            = errors.New("hash provided is invalid")
	Err_User_Present            = errors.New("user already present")
	Err_IncorrectUserOrPassword = errors.New("user or password provided was incorrect")
)

func NewUserService(User repository.UserRepo, utsi UserTokenInterface) *UserService {

	return &UserService{User, utsi}
}

func (us *UserService) LoginUser(email string, password string) (string, error) {

	userobject, err := us.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = utils.ComparePassword(userobject.Password, password)
	if err != nil {
		return "", Err_IncorrectUserOrPassword
	}
	return us.UserTokenService.GenerateAndSaveUserToken(userobject.ID)

}

func (us *UserService) SaveUser(email string, name string, password string) (string, error) {

	_, err := us.UserRepo.GetUserByEmail(email)

	if err == nil {
		return "", Err_User_Present
	}

	if err != repository.Err_UserNotFound {
		return "", err
	}

	timestamp := time.Now().Unix()
	encryptedPassword, err := utils.GenerateHashForPassword(password)
	if err != nil {
		return "", err
	}
	userObject := domain.NewUser(email, name, encryptedPassword, timestamp)
	err = us.UserRepo.SaveUser(userObject)
	if err != nil {
		return "", err
	}
	return us.UserTokenService.GenerateAndSaveUserToken(userObject.ID)
}

func (us *UserService) GetUserByID(id string) (*domain.User, error) {
	user, err := us.UserRepo.GetUserByID(id)

	return user, err
}
func (us *UserService) GetUserNamesByIDs(ids []string) ([]string, error) {
	user, err := us.UserRepo.GetUserNamesByIDs(ids)

	return user, err
}
func (us *UserService) CountUsersFromIDs(user_ids []string) (int64, error) {
	user, err := us.UserRepo.CountUsersFromIDs(user_ids)

	return user, err
}
