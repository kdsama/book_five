package service

import (
	"errors"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

var (
	MAX_LIST_COUNT = int64(5)
	MAX_LIST_SIZE  = int64(5)
)

var (
	err_ListCannotBeCreated    = errors.New("there was an issue while creating the list")
	err_ListCreationNotAllowed = errors.New("Error list creation is not allowed")
	err_ListSizeExceeded       = errors.New("List size exceeds the maximum size")
	err_CannotComment          = errors.New("user is not allowed to comment here.")
)

type UserListService struct {
	book          BookServiceInterface
	user          UserServiceInterface
	user_activity UserActivityServiceInterface
	comment       ListCommentServiceInterface
	userlistRepo  repository.UserListRepo
}

func NewUserListService(user UserServiceInterface, book BookServiceInterface, user_activity UserActivityServiceInterface, list_comment ListCommentServiceInterface, userlistRepo repository.UserListRepo) *UserListService {

	return &UserListService{book, user, user_activity, list_comment, userlistRepo}
}

func (uls *UserListService) SaveUserList(user_id string, about string, list_name string, book_ids []string) error {

	user, err := uls.user.GetUserByID(user_id)
	if err != nil {
		// User just might not be present
		// this check probably will be done on middleware as well , the jwt would be checked here
		return err
	}
	if len(book_ids) > int(MAX_LIST_SIZE) {
		return err_ListSizeExceeded
	}

	//remove duplicates
	var book_mapping map[string]int
	new_book_ids := []string{}
	for i := range book_ids {
		if _, ok := book_mapping[book_ids[i]]; !ok {
			new_book_ids = append(new_book_ids, book_ids[i])
		}
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

	userListObject := domain.NewUserList(user.ID, about, new_book_ids, list_name, timestamp)
	list_count, err := uls.CountExistingListsOfAUser(user_id)
	if err != nil {
		return err
	}
	// if one more added , we have to check whether the limit is reached or exceeded .
	if list_count+1 >= MAX_LIST_COUNT {
		return err_ListCreationNotAllowed
	}
	err = uls.userlistRepo.SaveUserList(userListObject)
	// no need to create an activity for creating the user list ???
	// for now lets skip it .
	return err
}

func (uls *UserListService) CountExistingListsOfAUser(user_id string) (int64, error) {

	count, err := uls.userlistRepo.CountExistingListsOfAUser(user_id)
	if err != nil && err != repository.Err_NoUserListFound {
		return 0, err
	}
	return count, nil

}

func (uls *UserListService) SaveComment(list_id string, user_id string, comment string) (string, error) {

	// check if list exists
	_, err := uls.userlistRepo.GetListByID(list_id)
	if err != nil {
		return "", err
	}
	// false check if user is allowed to do put a comment
	canComment := true

	if !canComment {
		return "", err_CannotComment
	}
	return uls.comment.SaveListComment(user_id, list_id, comment)
	//

}
