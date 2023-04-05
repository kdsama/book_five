package service

import (
	"errors"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

var (
	Err_MaliciousIntent = errors.New("this token doesnot belong to the user but was saved here. ")
	Err_TokenExpired    = errors.New("token was expired")
)

type UserTokenService struct {
	UserTokenRepo repository.UserTokenRepository
}

func NewUserTokenService(repo repository.UserTokenRepository) *UserTokenService {
	return &UserTokenService{repo}
}

func (uts *UserTokenService) GenerateUserToken(user_id string) (string, error) {
	return utils.CreateJWTToken(user_id)
}

// generate saves and returns the token
func (uts *UserTokenService) GenerateAndSaveUserToken(user_id string) (string, error) {

	timestamp := utils.GetCurrentTimestamp()
	// Generate JWT Token
	token, err := uts.GenerateUserToken(user_id)
	if err != nil {
		return "", err
	}
	userTokenObject := domain.NewUserToken(user_id, token, timestamp)
	err = uts.UserTokenRepo.SaveUserToken(userTokenObject)
	if err != nil {
		return "", err
	}
	return userTokenObject.Token, nil
}

func (uts *UserTokenService) GetUserTokenByID(user_id string) (*domain.UserToken, error) {
	return uts.UserTokenRepo.GetUserTokenByID(user_id)
}
func (uts *UserTokenService) GetUserByToken(token string) (*domain.UserToken, error) {
	return uts.UserTokenRepo.GetUserByToken(token)
}

func (uts *UserTokenService) ValidateUserTokenAndGetUserID(token string) (string, error) {

	id, err := utils.VerifyJWTToken(token)
	if err != nil {
		return "", err
	}
	usertoken, err := uts.GetUserByToken(token)
	if err != nil {
		return "", err
	}
	if id != usertoken.User_ID {
		return "", Err_MaliciousIntent
	}
	if usertoken.UpdatedAt >= utils.GetCurrentTimestamp() {
		return "", Err_TokenExpired
	}
	return id, nil
}
