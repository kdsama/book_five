package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kdsama/book_five/service"
)

type InputBook struct {
	Name          string   `json:"name"`
	Authors       []string `json:"authors"`
	Co_Authors    []string `json:"co_authors"`
	AudiobookUrls []string `json:"audiobook_urls"`
	EbookUrls     []string `json:"ebook_urls"`
	Hardcopies    []string `json:"hardcopies"`
	Categories    []string `json:"categories"`
}

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(bookservice service.BookService) *BookHandler {

	return &BookHandler{bookservice}

}

func (bh *BookHandler) Req(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		bh.postBook(w, req)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintln(http.StatusNotImplemented)))
	}
}

func (bh *BookHandler) postBook(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t InputBook
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}

	if t.Name == "" || len(t.Authors) == 0 || len(t.Categories) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	bh.service.SaveBook(t.Name, t.Authors, t.Co_Authors, t.AudiobookUrls, t.EbookUrls, t.Hardcopies, t.Categories)
}
