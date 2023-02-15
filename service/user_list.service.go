package service

import (
	"errors"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

var (
	MAX_LIST_COUNT = 5
)

var (
	err_ListCannotBeCreated    = errors.New("there was an issue while creating the list")
	err_ListCreationNotAllowed = errors.New("Error list creation is not allowed")
)

type UserListService struct {
	book         BookServiceInterface
	user         UserServiceInterface
	userlistRepo repository.UserListRepo
}

func NewUserListService(user UserServiceInterface, book BookServiceInterface, userlistRepo repository.UserListRepo) *UserListService {

	return &UserListService{book, user, userlistRepo}
}

func (uls *UserListService) SaveUserList(user_id string, list_name string, book_ids []string) error {

	user, err := uls.user.GetUserById(user_id)
	if err != nil {
		// User just might not be present
		return err
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

	userListObject := domain.NewUserList(user.Id, book_ids, list_name, timestamp)
	list_count, err := uls.CountExistingListsOfAUser(user.Id)
	if err != nil {
		return err
	}
	if list_count >= MAX_LIST_COUNT {
		return err_ListCreationNotAllowed
	}
	err = uls.userlistRepo.SaveUserList(userListObject)
	return err
}

func (uls *UserListService) CountExistingListsOfAUser(user_id string) (int, error) {

	count, err := uls.userlistRepo.CountExistingListsOfAUser(user_id)
	if err != nil && err != repository.Err_NoUserListFound {
		return 0, err
	}
	return count, nil

}
func countErrorsFromSlice(err_slice []error) int {
	counter := 0
	for i := range err_slice {
		if err_slice[i] != nil {
			counter++
		}
	}
	return counter
}
