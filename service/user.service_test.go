package service

import (
	"errors"
	"testing"
	"time"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type mockUserRepo struct {
	err error
}

var mockUsers []domain.User

func (mub *mockUserRepo) SaveUser(user *domain.User) error {
	mockUsers = append(mockUsers, *user)

	return mub.err
}
func (mub *mockUserRepo) GetUserByEmail(email string) (*domain.User, error) {
	if mub.err != nil {
		return &domain.User{}, mub.err
	}
	for i := range mockUsers {
		if mockUsers[i].Email == email {
			return &mockUsers[i], nil
		}

	}
	return &domain.User{}, repository.Err_UserNotFound
}
func TestSaveUser(t *testing.T) {
	TestSaveUserErrors(t)
	type userInput struct {
		email    string
		password string
	}
	type test struct {
		title string
		want  error
		input userInput
	}
	mockUsers = append(mockUsers, *domain.NewUser("already@exists.com", "SomepasswordNotRelevantInOurCase", time.Now().Unix()))

	userrepo := repository.NewUserRepository(&mockUserRepo{})
	userservice := NewUserService(*userrepo)

	testObject := &test{
		title: "should return nil as all conditions are correct",
		want:  nil,
		input: userInput{email: "doesnot@exist.com",
			password: "e86f78a8a3caf0b60d8e74e5942aa6d86dc150cd3c03338aef25b7d2d7e3acc7",
		},
	}
	got := userservice.SaveUser(testObject.input.email, testObject.input.password)
	if testObject.want != got {
		t.Errorf("wanted error %v but got %v", testObject.want, got)
	}
}

func TestSaveUserErrors(t *testing.T) {
	type userInput struct {
		email    string
		password string
	}
	type test struct {
		title   string
		want    error
		input   userInput
		check   string
		service *UserService
	}
	mockUsers = append(mockUsers, *domain.NewUser("already@exists.com", "SomepasswordNotRelevantInOurCase", time.Now().Unix()))
	userrepo := repository.NewUserRepository(&mockUserRepo{})
	userservice := NewUserService(*userrepo)
	userrepo1 := repository.NewUserRepository(&mockUserRepo{errors.New("Some error")})
	userservices1 := NewUserService(*userrepo1)
	table := []test{
		{title: "Should return error if user already exists",
			want: Err_User_Present,
			input: userInput{email: "already@exists.com",
				password: "e86f78a8a3caf0b60d8e74e5942aa6d86dc150cd3c03338aef25b7d2d7e3acc7",
			},
			check:   "!=",
			service: userservice,
		},
		{title: "If database returns an error that is not noRecordFoundError. ",
			want: Err_User_Present,
			input: userInput{email: "already@exists.com",
				password: "e86f78a8a3caf0b60d8e74e5942aa6d86dc150cd3c03338aef25b7d2d7e3acc7"},
			check:   "=",
			service: userservices1,
		}}

	for _, testObject := range table {
		t.Run(testObject.title, func(t *testing.T) {
			want := testObject.want

			got := testObject.service.SaveUser(testObject.input.email, testObject.input.password)
			switch testObject.check {
			case "!=":
				if got != want {
					t.Errorf("wanted error %v but got %v", want, got)
				}
			case "=":
				if got == want {
					t.Errorf("didnot want  error %v but got it", got)
				}

			}

		})
	}
}
