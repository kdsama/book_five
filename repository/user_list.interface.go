package repository

import "github.com/kdsama/book_five/domain"

type UserListRepo interface {
	SaveUserList(*domain.UserList) error
	CountExistingListsOfAUser(user_id string) (int, error)
}

type UserListRepository struct {
	UserListRepo
}

func NewUserListRepository(br UserListRepo) *UserListRepository {
	return &UserListRepository{br}
}
