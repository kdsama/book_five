package service

type UserServiceInterface interface {
	SaveUser(string, string) error
	CountUsersFromIDs([]string) (int64, error)
}

type UserDI struct {
	UserServiceInterface
}

func NewUserServiceInterface(br UserServiceInterface) *UserDI {
	return &UserDI{br}
}
