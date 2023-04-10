package repository

import "github.com/kdsama/book_five/domain"

type ListCommentRepo interface {
	SaveListComment(*domain.ListComment) error
	GetCommentsByListID(list_id string) ([]domain.ListComment, error)
}

type ListCommentRepository struct {
	ListCommentRepo
}

func NewListCommentRepository(br ListCommentRepo) *ListCommentRepository {
	return &ListCommentRepository{br}
}
