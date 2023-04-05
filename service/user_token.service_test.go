package service

import domain "github.com/kdsama/book_five/domain"

type MockUserTokenService struct{}

func (muts *MockUserTokenService) GenerateUserToken(user_id string) (string, error) { return "", nil }
func (muts *MockUserTokenService) GenerateAndSaveUserToken(user_id string) (string, error) {
	return "", nil
}
func (muts *MockUserTokenService) GetUserTokenByID(user_id string) (*domain.UserToken, error) {
	return &domain.UserToken{}, nil
}
func (muts *MockUserTokenService) ValidateUserTokenAndGetUserID(token string) (string, error) {
	return "", nil
}
