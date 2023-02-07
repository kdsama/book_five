package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kdsama/book_five/service"
)

type MockUserService struct {
	err error
}

func (mbs *MockUserService) SaveUser(email string, password string) error {
	return mbs.err
}
func TestUser(t *testing.T) {

	TestPostUser(t)

}

func TestPostUser(t *testing.T) {
	TestPostUserErrors(t)
	type requestContent struct {
		bdy  map[string]interface{}
		want int
	}
	us := service.NewUserServiceInterface(&MockUserService{})
	userHandler := NewUserHandler(*us)

	t.Run("Correct User and PW", func(t *testing.T) {
		request := &requestContent{
			bdy: map[string]interface{}{"email": "kd@correctpassword.com",
				"wwl": "aaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbb!3457gsdrwerwera123"},
			want: http.StatusOK}

		body, _ := json.Marshal(request.bdy)
		req, err := http.NewRequest("POST", "/api/v1/book", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Req)

		handler.ServeHTTP(rr, req)
		want := request.want
		if got := rr.Code; got != want {
			t.Errorf("handler returned wrong status code: got %v want %v",
				got, want)
		}
	})
}
func TestPostUserErrors(t *testing.T) {

	us := service.NewUserServiceInterface(&MockUserService{errors.New("Some error")})
	userHandler := NewUserHandler(*us)

	type requestContent struct {
		bdy   map[string]interface{}
		want  int
		title string
	}
	requests := []requestContent{

		{bdy: map[string]interface{}{
			"email": "kd@nopassword.com",
			"wwl":   "",
		},
			want:  http.StatusBadRequest,
			title: "empty password",
		},
		{bdy: map[string]interface{}{
			"email": "",
			"wwl":   "passwordEmptyEmail",
		},
			want:  http.StatusBadRequest,
			title: "empty email",
		},
		{bdy: map[string]interface{}{
			"email": "Incorrect Email format",
			"wwl":   "passwordEmptyEmail",
		},
			want:  http.StatusBadRequest,
			title: "Incorrect EmailFormat",
		},
		{bdy: map[string]interface{}{
			"email": "kd@smallpw.com",
			"wwl":   "nonEmptybutsmallpw",
		},
			want:  http.StatusBadRequest,
			title: "Password Length is too Small. Preface :- encrypted password is supposed to be shared",
		},
	}

	for _, request := range requests {
		t.Run(request.title, func(t *testing.T) {
			body, _ := json.Marshal(request.bdy)
			req, err := http.NewRequest("POST", "/api/v1/book", bytes.NewReader(body))
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(userHandler.Req)

			handler.ServeHTTP(rr, req)
			want := request.want
			if got := rr.Code; got != want {
				t.Errorf("handler returned wrong status code: got %v want %v",
					got, want)
			}

		})
	}
}
