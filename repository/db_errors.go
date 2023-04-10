package repository

import "errors"

var (
	Err_UserNotFound          = errors.New("user couldnot be found")
	Err_NoUserListFound       = errors.New("No user list was found")
	Err_BookNotFound          = errors.New("book couldnot be found ")
	Err_NoBooksInCategory     = errors.New("no Book is present in this category")
	ErrWriteRecord            = errors.New("cannot write to repository")
	Err_CategoryNotFound      = errors.New("category couldnot be found ")
	Err_NoCategorysInCategory = errors.New("no categories are present in this category")
	Err_UserTokenNotFound     = errors.New("user token couldnot be found")
	Err_UserListNotFound      = errors.New("user list couldnot be found")
	Err_CannotLoadComments    = errors.New("cannot load comments, please try again later")
)
