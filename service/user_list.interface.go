package service

import (
	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/entity"
)

type UserListServiceInterface interface {
	SaveUserList(user_id string, about string, list_name string, book_ids []string) error
	CountExistingListsOfAUser(user_id string) (int64, error)
	SaveComment(list_id string, user_id string, comment string) (string, error)
	GetComments(list_id string) ([]domain.ListComment, error)
	React(user_id string, list_id string, reaction string, comment_id string) error
	UpdateUserListReactions(list_id string, reaction entity.Reaction) error
}

type UserListDI struct {
	UserListServiceInterface
}

func NewUserListServiceInterface(br UserListServiceInterface) *UserListDI {
	return &UserListDI{br}
}
