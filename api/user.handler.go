package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/kdsama/book_five/service"
)

type InputUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"wwl"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"wwl"`
}

type UserHandler struct {
	service service.UserDI
}

var MIN_PASSWORD_HASH_LENGTH int = 30

func NewUserHandler(bookservice service.UserDI) *UserHandler {

	return &UserHandler{bookservice}

}

func (bh *UserHandler) Req(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		bh.postUser(w, req)
	case http.MethodGet:
		bh.loginUser(w, req)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintln(http.StatusNotImplemented)))
	}
}

func (bh *UserHandler) loginUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t InputUser
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}
	ok := validateUserInformation(t)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	token, err := bh.service.LoginUser(t.Email, t.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln(err)))
		return
	}
	// Set the JWT token in a cookie
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("ok"))
}

func (bh *UserHandler) postUser(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var t InputUser
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}

	ok := validateUserInformation(t)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	err = bh.service.SaveUser(t.Email, t.Name, t.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln(http.StatusInternalServerError)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("ok"))
}

func validateUserInformation(t InputUser) bool {
	if t.Email == "" || t.Password == "" {

		return false
	}
	// Check if incorrect email
	_, err := mail.ParseAddress(t.Email)
	if err != nil {
		return false
	}
	if len(t.Password) < MIN_PASSWORD_HASH_LENGTH {
		return false
	}

	return true
}
