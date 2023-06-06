package repository

import (
	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/entity"
)

type UserListRepo interface {
	SaveUserList(*domain.UserList) error
	CountExistingListsOfAUser(user_id string) (int64, error)
	GetListByID(list_id string) (*domain.UserList, error)
	UpdateUserListReactions(list_id string, reaction entity.Reaction) error
}

type UserListRepository struct {
	UserListRepo
}

func NewUserListRepository(br UserListRepo) *UserListRepository {
	return &UserListRepository{br}
}
