package service

import (
	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type MockUserTokenService struct{ err error }

var mockTokens []domain.UserToken

func (muts *MockUserTokenService) GenerateUserToken(user_id string) (string, error) { return "", nil }
func (muts *MockUserTokenService) GenerateAndSaveUserToken(user_id string) (string, error) {
	token, err := utils.CreateJWTToken(user_id)
	if err != nil {
		return "", err
	}
	mockTokens = append(mockTokens, domain.UserToken{User_ID: user_id, Token: token})
	return token, nil
}
func (muts *MockUserTokenService) GetUserTokenByID(user_id string) (*domain.UserToken, error) {
	if muts.err != nil {
		return &domain.UserToken{}, muts.err
	}
	for i := range mockTokens {
		if mockTokens[i].User_ID == user_id {
			return &mockTokens[i], nil
		}
	}
	return &domain.UserToken{}, repository.Err_UserTokenNotFound
}
func (muts *MockUserTokenService) ValidateUserTokenAndGetUserID(token string) (string, error) {
	return "", nil
}
