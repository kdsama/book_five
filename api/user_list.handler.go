package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kdsama/book_five/service"
)

type InputUserList struct {
	User_id   string   `json:"user_id"`
	About     string   `json:"about"`
	List_name string   `json:"list_name"`
	Book_ids  []string `json:"book_ids"`
}

type UserListHandler struct {
	service service.UserListDI
}

func NewUserListHandler(bookservice service.UserListDI) *UserListHandler {

	return &UserListHandler{bookservice}

}

func (bh *UserListHandler) Req(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		bh.postUserList(w, req)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintln(http.StatusNotImplemented)))
	}
}

func (bh *UserListHandler) postUserList(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var t InputUserList
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}

	ok := validatePostUserList(t)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	err = bh.service.SaveUserList(t.User_id, t.About, t.List_name, t.Book_ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln(http.StatusInternalServerError)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("ok"))
}

func validatePostUserList(t InputUserList) bool {
	if t.User_id == "" || t.About == "" || len(t.List_name) == 0 {
		// the length of books should be atleast 1 right ?
		return false
	}

	return true
}
