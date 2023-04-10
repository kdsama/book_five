package service

import (
	"time"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type ListCommentService struct {
	listcommentRepo repository.ListCommentRepository
}

func NewListCommentService(ListComment repository.ListCommentRepository) *ListCommentService {
	return &ListCommentService{ListComment}
}
func (bs *ListCommentService) SaveListComment(list_id string, user_id string, comment string) (string, error) {
	// validation checks already done in the http handler .
	// We just need to save it for now.
	timestamp := time.Now().Unix()

	listObject := domain.NewListComment(list_id, user_id, comment, timestamp)
	return listObject.ID, bs.listcommentRepo.SaveListComment(listObject)

}
