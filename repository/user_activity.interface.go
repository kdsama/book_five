package repository

import "github.com/kdsama/book_five/domain"

type UserActivityRepo interface {
	SaveUserActivity(*domain.UserActivity) error
	UpdateUserActivty(*domain.UserActivity) error
	GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error)
	GetUserReactionActivityByUserAndListID(user_id string, list_id string) (*domain.UserActivity, error)
}

type UserActivityRepository struct {
	UserActivityRepo
}

func NewUserActivityRepository(br UserActivityRepo) *UserActivityRepository {
	return &UserActivityRepository{br}
}
