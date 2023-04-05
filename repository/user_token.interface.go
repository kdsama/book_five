package repository

import "github.com/kdsama/book_five/domain"

type UserTokenRepo interface {
	SaveUserToken(*domain.UserToken) error
	GetUserTokenByID(string) (*domain.UserToken, error)
	GetUserByToken(string) (*domain.UserToken, error)
}

type UserTokenRepository struct {
	UserTokenRepo
}

func NewUserTokenRepository(br UserTokenRepo) *UserTokenRepository {
	return &UserTokenRepository{br}
}
