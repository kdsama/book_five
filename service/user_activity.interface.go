package service

import (
	"github.com/kdsama/book_five/domain"
)

type UserActivityServiceInterface interface {
	SaveUserActivity(user_id string, action string, reaction string, receiver string, list_id string, comment_id string, review_id string) error
	UpdateUserActivty(user_id string, action string, reaction string, receiver string, list_id string, comment_id string, review_id string) error
	GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error)
	GetUserReactionActivityByUserAndListID(user_id string, list_id string) (*domain.UserActivity, error)
}

type UserActivityDI struct {
	UserActivityServiceInterface
}

func NewUserActivityServiceInterface(br UserActivityServiceInterface) *UserActivityDI {
	return &UserActivityDI{br}
}
