package service

type UserListServiceInterface interface {
	SaveUserList(user_id string, about string, list_name string, book_ids []string) error
	CountExistingListsOfAUser(user_id string) (int, error)
}

type UserListDI struct {
	UserListServiceInterface
}

func NewUserListServiceInterface(br UserListServiceInterface) *UserListDI {
	return &UserListDI{br}
}
