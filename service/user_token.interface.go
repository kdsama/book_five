package service

import "github.com/kdsama/book_five/domain"

type UserTokenInterface interface {
	GenerateUserToken(user_id string) (string, error)
	GenerateAndSaveUserToken(user_id string) (string, error)
	GetUserTokenByID(user_id string) (*domain.UserToken, error)
	ValidateUserTokenAndGetUserID(token string) (string, error)
}

type UserTokenDI struct {
	UserTokenInterface
}

func NewUserTokenServiceInterface(uti UserTokenInterface) *UserTokenDI {
	return &UserTokenDI{UserTokenInterface: uti}

}
