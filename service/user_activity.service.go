package service

import (
	"strings"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type UserActivityService struct {
	user UserServiceInterface

	userActivityRepo repository.UserActivityRepository
}

func NewUserActivityService(user UserServiceInterface, userActivityRepo repository.UserActivityRepository) *UserActivityService {

	return &UserActivityService{user, userActivityRepo}
}

func (uls *UserActivityService) SaveUserActivity(user_id string, action string, receiver string, list_id string, comment_id string, review_id string) error {

	user_names, err := uls.user.GetUserNamesByIDs([]string{receiver})
	if err != nil {
		// User just might not be present
		return err
	}
	if len(user_names) != 1 {
		// one of the user is not present it seems
		return repository.Err_UserNotFound
	}

	timestamp := utils.GetCurrentTimestamp()
	desc := "$1 $2 on your $3 "
	if user_id == receiver {
		desc = strings.Replace(desc, "$1", "You", 1)

	} else {
		desc = strings.Replace(desc, "$1", user_names[1], 1)
	}
	if action == "comment" {
		desc = strings.Replace(desc, "$2", "commented", 1)
	} else if action == "reaction" {
		desc = strings.Replace(desc, "$2", "reacted", 1)
	}
	if comment_id != "" {
		desc = strings.Replace(desc, "$2", "Comment", 1)
	} else if list_id != "" {
		desc = strings.Replace(desc, "$2", "List", 1)
	} else if review_id != "" {
		desc = strings.Replace(desc, "$2", "Book Review", 1)
	}
	UserActivityObject := domain.NewUserActivity(user_id, action, receiver, list_id, comment_id, review_id, desc, timestamp)
	err = uls.userActivityRepo.SaveUserActivity(UserActivityObject)
	return err
}

func (uls *UserActivityService) GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error) {
	return uls.userActivityRepo.GetLastUserActivityByUserID(user_id)
}
