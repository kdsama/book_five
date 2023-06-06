package service

import (
	"testing"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/utils"
)

type MockUserActivityService struct {
	err error
}

var mockactivities = []domain.UserActivity{}

func (mua *MockUserActivityService) SaveUserActivity(user_id string, action string, reaction string, receiver string, list_id string, comment_id string, review_id string) error {
	mockactivities = append(mockactivities, *domain.NewUserActivity(user_id, action, reaction, receiver, list_id, comment_id, review_id, "Dont care here", utils.GetCurrentTimestamp()))
	return nil
}

func (mua *MockUserActivityService) UpdateUserActivty(user_id string, action string, reaction string, receiver string, list_id string, comment_id string, review_id string) error {
	if mua.err != nil {
		return mua.err
	}
	for i := range mockactivities {
		if mockactivities[i].List_ID == list_id && mockactivities[i].User_ID == user_id && mockactivities[i].Receiver == receiver {
			mockactivities[i].Action = action
			mockactivities[i].Reaction = reaction
		}

	}
	return nil
}

func (mua *MockUserActivityService) GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error) {
	return &mockactivities[len(mockactivities)-1], nil
}

func (mua *MockUserActivityService) GetUserReactionActivityByUserAndListID(user_id string, list_id string) (*domain.UserActivity, error) {
	if mua.err != nil {
		return &domain.UserActivity{}, mua.err
	}
	for i := range mockactivities {
		if mockactivities[i].List_ID == list_id && mockactivities[i].Action == "reaction" {
			return &mockactivities[i], nil
		}

	}
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
