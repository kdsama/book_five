package service

import "github.com/kdsama/book_five/domain"

type UserListServiceInterface interface {
	SaveUserList(user_id string, about string, list_name string, book_ids []string) error
	CountExistingListsOfAUser(user_id string) (int64, error)
	SaveComment(list_id string, user_id string, comment string) (string, error)
	GetComments(list_id string) ([]domain.ListComment, error)
}

type UserListDI struct {
	UserListServiceInterface
}

func NewUserListServiceInterface(br UserListServiceInterface) *UserListDI {
	return &UserListDI{br}
}
