package service

import (
	"strings"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type UserActivityService struct {
	user             UserServiceInterface
	userActivityRepo repository.UserActivityRepo
	userListRepo     repository.UserListRepo
}

func NewUserActivityService(user UserServiceInterface, book repository.UserActivityRepo, userlistRepo repository.UserListRepo) *UserActivityService {

	return &UserActivityService{user, book, userlistRepo}
}

func (uls *UserActivityService) SaveUserActivity(user_id string, action string, receiver string, list_id string, comment_id string, review_id string) error {

	user_count, err := uls.user.GetUserNamesByIDs([]string{user_id, receiver})
	if err != nil {
		// User just might not be present
		return err
	}
	if user_count == 0 {
		return repository.Err_UserNotFound
	}
	// No Comment as it is a new List
	// No need to check if the book exist
	// We will save the book separately first and only then pass it to the user list
	// If they dont we need to create the books probably
	// Make them unverified. The books needs to be verified
	// Make a parameter , if that parameter is true, The user can create more than 5 lists,
	// lets say 20 lists. That parameter reveals for the particular user how many books
	// can he or she or they can add.
	// In the future it will be related to some kind of book score for that particular user

	timestamp := utils.GetCurrentTimestamp()
	desc := "$1 $2 on your $3 "
	if user_id == receiver {
		desc = strings.Replace(desc, "$1", "You", 1)

	} else {
		desc = strings.Replace(desc, "$2", "", 1)
		if action == "comment" {
			desc += " commented"
		} else {
			desc += " reacted"
		}

	}
	userListObject := domain.NewUserActivity(user_id, action, list_id, comment_id, review_id, desc, timestamp)
	list_count, err := uls.CountExistingListsOfAUser(user_id)
	if err != nil {
		return err
	}
	if list_count >= MAX_LIST_COUNT {
		return err_ListCreationNotAllowed
	}
	err = uls.userActivityRepo.SaveUserActivity(userListObject)
	return err
}

func (uls *UserActivityService) CountExistingListsOfAUser(user_id string) (int, error) {

	count, err := uls.userListRepo.CountExistingListsOfAUser(user_id)
	if err != nil {
		return 0, err
	}
	return count, nil

}
