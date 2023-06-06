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

func (uas *UserActivityService) SaveUserActivity(user_id string, action string, reaction string, receiver string, list_id string, comment_id string, review_id string) error {

	user_names, err := uas.user.GetUserNamesByIDs([]string{receiver})
	if err != nil {
		// User just might not be present
		return err
	}

	if len(user_names) != 1 {
		// one of the user is not present it seems
		return repository.Err_UserNotFound
	}

	timestamp := utils.GetCurrentTimestamp()
	desc := GenerateDescription(user_id, receiver, user_names, action, comment_id, list_id, review_id)
	UserActivityObject := domain.NewUserActivity(user_id, action, reaction, receiver, list_id, comment_id, review_id, desc, timestamp)
	err = uas.userActivityRepo.SaveUserActivity(UserActivityObject)
	return err
}

func GenerateDescription(user_id string, receiver string, user_names []string, action string, comment_id string, list_id string, review_id string) string {
	desc := "$1 $2 on your $3 "
	if user_id == receiver {
		desc = strings.Replace(desc, "$1", "You", 1)

	} else {
		desc = strings.Replace(desc, "$1", user_names[0], 1)
	}
	if action == "comment" {
		desc = strings.Replace(desc, "$2", "commented", 1)
	} else if action == "reaction" {
		desc = strings.Replace(desc, "$2", "reacted", 1)
	}
	if comment_id != "" {
		desc = strings.Replace(desc, "$3", "Comment", 1)

	} else if list_id != "" {
		desc = strings.Replace(desc, "$3", "List", 1)
	} else if review_id != "" {
		desc = strings.Replace(desc, "$3", "Book Review", 1)
	}
	return desc
}

func (uas *UserActivityService) GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error) {
	return uas.userActivityRepo.GetLastUserActivityByUserID(user_id)
}

func (uas *UserActivityService) GetUserReactionActivityByUserAndListID(user_id string, list_id string) (*domain.UserActivity, error) {
	return uas.userActivityRepo.GetUserReactionActivityByUserAndListID(user_id, list_id)
}

func (uas *UserActivityService) UpdateUserActivty(user_id string, action string, reaction string, receiver string, list_id string, comment_id string, review_id string) error {

	user_names, err := uas.user.GetUserNamesByIDs([]string{receiver})
	if err != nil {
		// User just might not be present
		return err
	}

	if len(user_names) != 1 {
		// one of the user is not present it seems
		return repository.Err_UserNotFound
	}

	timestamp := utils.GetCurrentTimestamp()

	UserActivityObject := domain.NewUserActivity(user_id, action, reaction, receiver, list_id, comment_id, review_id, "", timestamp)
	return uas.userActivityRepo.UpdateUserActivty(UserActivityObject)
}
