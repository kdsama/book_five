package service

type ListCommentServiceInterface interface {
	SaveListComment(user_id string, list_id string, comment string) (string, error)
}

type ListCommentDI struct {
	ListCommentServiceInterface
}

func NewListCommentServiceInterface(br ListCommentServiceInterface) *ListCommentDI {
	return &ListCommentDI{br}
}
