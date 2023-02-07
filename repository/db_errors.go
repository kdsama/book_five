package repository

import "errors"

var (
	Err_UserNotFound = errors.New("user couldnot be found")

	Err_BookNotFound          = errors.New("book couldnot be found ")
	Err_NoBooksInCategory     = errors.New("no Book is present in this category")
	ErrWriteRecord            = errors.New("cannot write to repository")
	Err_CategoryNotFound      = errors.New("category couldnot be found ")
	Err_NoCategorysInCategory = errors.New("no categories are present in this category")
)