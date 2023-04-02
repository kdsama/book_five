package service

import (
	"testing"

	domain "github.com/kdsama/book_five/domain"
)

type MockUserActivityService struct {
	err error
}

func (mua *MockUserActivityService) SaveUserActivity(user_id string, action string, receiver string, list_id string, comment_id string, review_id string) error {
	return nil
}

func (mua *MockUserActivityService) GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error) {
	return &domain.UserActivity{}, nil
}
func TestUserActivity(t *testing.T) {
	TestSaveUserActivity(t)
}

func TestSaveUserActivity(t *testing.T) {
	TestSaveUserActivityErrors(t)
}
func TestSaveUserActivityErrors(t *testing.T) {

	//Users DoNot Exist
	//List is passed but it donot exist
	//Comment is passed but it donot exist
	//List and Comment passed but both dono exist

}
