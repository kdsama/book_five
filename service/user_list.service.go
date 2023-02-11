package service

import (
	"errors"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

var (
	err_ListCannotBeCreated = errors.New("there was an issue while creating the list")
)

type UserListService struct {
	book         BookServiceInterface
	user         UserServiceInterface
	userlistRepo repository.UserListRepo
}

func NewUserListService(user UserServiceInterface, book BookServiceInterface, userlistRepo repository.UserListRepo) *UserListService {

	return &UserListService{book, user, userlistRepo}
}

func (uls *UserListService) SaveUserList(user_id string, list_name string, books []domain.Book) error {

	user, err := uls.user.GetUserById(user_id)
	if err != nil {
		// User just might not be present
		return err
	}

	// No Comment as it is a new List
	// So we need to check if the books exist
	// If they dont we need to create the books probably
	// Make them unverified. The books needs to be verified
	// Make a parameter , if that parameter is true, The user can create more than 5 lists,
	// lets say 20 lists. That parameter reveals for the particular user how many books
	// can he or she or they can add.
	// In the future it will be related to some kind of book score for that particular user
	//
	idSlice, errSlice, errorCount := uls.book.UpsertBooksAndGetIDs(books)

	if errorCount == len(idSlice) {
		return err_ListCannotBeCreated
	}

	timestamp := utils.GetCurrentTimestamp()

	userListObject := domain.NewUserList(user.Id, idSlice, list_name, timestamp)

}
