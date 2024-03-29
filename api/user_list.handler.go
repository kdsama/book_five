package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kdsama/book_five/api/middleware"
	"github.com/kdsama/book_five/service"
)

type InputUserList struct {
	User_ID  string   `json:"user_id"`
	About    string   `json:"about"`
	Name     string   `json:"list_name"`
	Book_ids []string `json:"book_ids"`
}

type InputListComment struct {
	User_ID string
	Comment string `json:"comment"`
	List_ID string `json:"list_id"`
}

type InputListReaction struct {
	User_ID    string
	Comment_ID string `json:"comment_id"`
	Reaction   string `json:"reaction"`
	List_ID    string `json:"list_id"`
}

type UserListHandler struct {
	service service.UserListDI
}

func NewUserListHandler(bookservice service.UserListDI) *UserListHandler {

	return &UserListHandler{bookservice}

}

func (bh *UserListHandler) Req(w http.ResponseWriter, req *http.Request) {
	// switch req.URL.Path{
	// 	fmt.Println
	// }
	fmt.Println(req.URL.Path)
	switch req.URL.Path {
	case "/api/v1/userlist":
		switch req.Method {
		case http.MethodPost:
			bh.postUserList(w, req)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(fmt.Sprintln(http.StatusNotImplemented)))
		}
	case "/api/v1/userlist/comment":
		switch req.Method {
		case http.MethodGet:
			bh.getListComments(w, req)
		case http.MethodPost:

			bh.postListComment(w, req)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(fmt.Sprintln(http.StatusNotImplemented)))
		}

	case "/api/v1/userlist/reaction":
		switch req.Method {
		case http.MethodPost:
			bh.postListReaction(w, req)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(fmt.Sprintln(http.StatusNotImplemented)))
		}
	}

}

func (u *UserListHandler) Handler() http.Handler {
	return http.HandlerFunc(u.Req)
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
	err = bh.service.SaveUserList(t.User_ID, t.About, t.Name, t.Book_ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("ok"))
}

func (bh *UserListHandler) postListComment(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t InputListComment
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}

	user_id := req.Context().Value(middleware.UserContextKey).(string)

	t.User_ID = user_id
	ok := validatePostListComment(t)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}

	comment_id, err := bh.service.SaveComment(t.List_ID, t.User_ID, t.Comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(comment_id))

}

func (bh *UserListHandler) postListReaction(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t InputListReaction
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}

	user_id := req.Context().Value(middleware.UserContextKey).(string)

	t.User_ID = user_id
	ok := validatePostListReaction(t)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}

	err = bh.service.React(t.User_ID, t.List_ID, t.Reaction, t.Comment_ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("ok"))

}

func (bh *UserListHandler) getListComments(w http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()
	list_id := queryValues.Get("list_id")
	if list_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	listcomments, err := bh.service.GetComments(list_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	listcommentJSON, err := json.Marshal(listcomments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(listcommentJSON)
}

func validatePostListComment(t InputListComment) bool {
	if t.User_ID == "" || t.Comment == "" || len(t.List_ID) == 0 {
		// the length of books should be atleast 1 right ?
		return false
	}

	return true
}

func validatePostListReaction(t InputListReaction) bool {
	if t.User_ID == "" || t.Reaction == "" || len(t.List_ID) == 0 {
		// the length of books should be atleast 1 right ?
		fmt.Println("Is here the oppression ? ", t)
		return false
	}

	return true
}

func validatePostUserList(t InputUserList) bool {
	if t.User_ID == "" || t.About == "" || len(t.Name) == 0 {
		// the length of books should be atleast 1 right ?
		return false
	}

	return true
}
