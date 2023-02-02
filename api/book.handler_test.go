package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kdsama/book_five/service"
)

type MockBookService struct {
}

func (mbs *MockBookService) SaveBook(name string, authors []string, co_authors []string, audiobook_urls []string,
	ebook_urls []string, hard_copies []string, categories []string) error {
	return nil
}
func TestBookService(t *testing.T) {

	// Test Save Book
	TestPostBook(t)
	// Test GetBook
	// TestGetBook(t)
}

func TestPostBook(t *testing.T) {

	TestPostBookErrors(t)
	r := service.NewBookServiceInterface(&MockBookService{})
	Bookhandler := NewBookHandler(*r)
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll

	// pass 'nil' as the third parameter.

	bdy := map[string]interface{}{
		"name":           "Never Finished",
		"authors":        []string{"David Goggins"},
		"categories":     []string{"Self-help"},
		"audiobook_urls": []string{"www.google.com"},
	}
	body, _ := json.Marshal(bdy)
	req, err := http.NewRequest("POST", "/api/v1/book", bytes.NewReader(body))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Bookhandler.Req)

	handler.ServeHTTP(rr, req)
	want := http.StatusOK
	if got := rr.Code; got != want {
		t.Errorf("handler returned wrong status code: got %v want %v",
			got, want)
	}

}

func TestPostBookErrors(t *testing.T) {
	r := service.NewBookServiceInterface(&MockBookService{})
	Bookhandler := NewBookHandler(*r)
	type requestContent struct {
		bdy   map[string]interface{}
		want  int
		title string
	}
	requests := []requestContent{

		{bdy: map[string]interface{}{
			"name":           "Never Finished",
			"authors":        []string{""},
			"categories":     []string{"Self-help"},
			"audiobook_urls": []string{"www.google.com"},
		},
			want:  http.StatusBadRequest,
			title: "Request without authors",
		},
		{bdy: map[string]interface{}{
			"name":           "Never Finished",
			"authors":        []string{""},
			"categories":     []string{"Self-help"},
			"audiobook_urls": []string{"www.google.com"},
		},
			want:  http.StatusBadRequest,
			title: "Request with an empty author ",
		},

		{bdy: map[string]interface{}{
			"name":           "Never Finished",
			"authors":        []string{"David Goggins"},
			"categories":     []string{},
			"audiobook_urls": []string{"www.google.com"},
		},
			want:  http.StatusBadRequest,
			title: "Request without Categories",
		},

		{bdy: map[string]interface{}{
			"name":           "Never Finished",
			"authors":        []string{"David Goggins"},
			"categories":     []string{""},
			"audiobook_urls": []string{"www.google.com"},
		},
			want:  http.StatusBadRequest,
			title: "Request with an empty category",
		},
		{bdy: map[string]interface{}{
			"inc1":       "Never Finished",
			"incorrect2": []string{"David Goggins"},
			"incor3":     []string{"self-help"},
			"none_here":  []string{"www.google.com"},
		},
			want:  http.StatusBadRequest,
			title: "Request with all wrong json keys ",
		},
		{bdy: map[string]interface{}{
			"name":       "Never Finished",
			"incorrect2": []string{"David Goggins"},
			"incor3":     []string{"self-help"},
			"none_here":  []string{"www.google.com"},
		},
			want:  http.StatusBadRequest,
			title: "Request with only One Correct json Key ",
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
			handler := http.HandlerFunc(Bookhandler.Req)

			handler.ServeHTTP(rr, req)
			want := request.want
			if got := rr.Code; got != want {
				t.Errorf("handler returned wrong status code: got %v want %v",
					got, want)
			}

		})
	}

}
