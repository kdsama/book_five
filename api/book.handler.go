package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kdsama/book_five/service"
)

type InputBook struct {
	Name          string   `json:"name"`
	Image_Url     string   `json:"image_url"`
	Authors       []string `json:"authors"`
	Co_Authors    []string `json:"co_authors"`
	AudiobookUrls []string `json:"audiobook_urls"`
	EbookUrls     []string `json:"ebook_urls"`
	Hardcopies    []string `json:"hardcopies"`
	Categories    []string `json:"categories"`
}

type BookHandler struct {
	service service.BookDI
}

func NewBookHandler(bookservice service.BookDI) *BookHandler {

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

	IsBadRequest := validatePostBook(t)
	if IsBadRequest {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	bh.service.SaveBook(t.Name, t.Image_Url, t.Authors, t.Co_Authors, t.AudiobookUrls, t.EbookUrls, t.Hardcopies, t.Categories)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("ok"))
}

func validatePostBook(t InputBook) bool {
	if t.Name == "" || len(t.Authors) == 0 || len(t.Categories) == 0 {

		return true
	}
	for i := range t.Authors {
		if strings.Trim(t.Authors[i], " ") == "" {
			return true
		}
	}
	for i := range t.Categories {
		if strings.Trim(t.Categories[i], " ") == "" {
			return true
		}
	}
	return false
}
