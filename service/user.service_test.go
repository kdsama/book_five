package service

import (
	"testing"
	"time"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type mockUserRepo struct {
}

func (mub *mockUserRepo) SaveUser(user *domain.User) error {
	return nil
}
func TestSaveUser(t *testing.T) {
	TestSaveUserErrors(t)
}

func TestSaveUserErrors(t *testing.T) {
	type userInput struct {
		email    string
		password string
	}
	type test struct {
		title string
		want  error
		input userInput
	}
	table := []test{{title: "Unencrypted password sent to service", want: Err_Invalid_Hash, input: userInput{email: "kd@gmail.com", password: "Something that is not a hash"}}}
	timestamp := time.Now().Unix()
	userrepo := repository.NewUserRepository(&mockUserRepo{})
	userservice := NewUserService(*userrepo)
	for _, testObject := range table {
		t.Run(testObject.title, func(t *testing.T) {
			want := testObject.want
			input := domain.NewUser(testObject.input.email, testObject.input.password, timestamp)
			got := userservice.SaveUser(input)
			if got != want {
				t.Errorf("wanted error %v but got %v", want, got)
			}
		})
	}
}
