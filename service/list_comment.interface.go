package service

import "github.com/kdsama/book_five/domain"

type ListCommentServiceInterface interface {
	SaveListComment(user_id string, list_id string, comment string) (string, error)
	GetCommentsByListID(list_id string) ([]domain.ListComment, error)
}

type ListCommentDI struct {
	ListCommentServiceInterface
}

func NewListCommentServiceInterface(br ListCommentServiceInterface) *ListCommentDI {
	return &ListCommentDI{br}
}
