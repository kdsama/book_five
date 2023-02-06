package service

import (
	"testing"

	"github.com/kdsama/book_five/domain"
)

type mockUserRepo struct {
}

func (mub *mockUserRepo) SaveUser(user *domain.User) error {
	return nil
}
func TestSaveUser(t *testing.T) {
	// TestSaveUserErrors(t)
}

// func TestSaveUserErrors(t *testing.T) {
// 	type userInput struct {
// 		email    string
// 		password string
// 	}
// 	type test struct {
// 		title string
// 		want  error
// 		input userInput
// 	}
// 	table := []test{{title: "Unencrypted password sent to service", want: Err_Invalid_Hash, input: userInput{email: "kd@gmail.com", password: "Something that is not a hash"}}}

// 	userrepo := repository.NewUserRepository(&mockUserRepo{})
// 	userservice := NewUserService(*userrepo)
// 	for _, testObject := range table {
// 		t.Run(testObject.title, func(t *testing.T) {
// 			want := testObject.want

// 			got := userservice.SaveUser(testObject.input.email, testObject.input.password)
// 			if got != want {
// 				t.Errorf("wanted error %v but got %v", want, got)
// 			}
// 		})
// 	}
// }
