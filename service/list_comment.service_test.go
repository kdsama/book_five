package service

import "github.com/kdsama/book_five/domain"

type mocklistcommentservice struct{ err error }

var mockListComment map[string][]domain.ListComment

func (mls *mocklistcommentservice) SaveListComment(user_id string, list_id string, comment string) (string, error) {
	if mls.err != nil {
		return "", mls.err
	}
	obj := *domain.NewListComment(list_id, user_id, comment, 0)
	mockListComment[list_id] = append(mockListComment[list_id], obj)
	return obj.ID, nil
}

func (mls *mocklistcommentservice) GetCommentsByListID(list_id string) ([]domain.ListComment, error) {

	return []domain.ListComment{}, nil
}
