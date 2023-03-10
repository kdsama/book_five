package service

import (
	"testing"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type MockBookRepository struct {
	types error
}

type MockCategoryService struct {
	saveError string
	getError  string
}

var MockBooks = []*domain.Book{}

func (mbr *MockBookRepository) SaveBook(NewBook *domain.Book) error {
	if mbr.types != nil {
		return mbr.types
	}
	MockBooks = append(MockBooks, NewBook)
	return nil
}

func (mcs *MockCategoryService) SaveCategory(name string, categories []string) error {
	if mcs.saveError != "" {
		return repository.ErrWriteRecord
	}
	return nil
	//
}

func (mcs *MockCategoryService) GetIdsByNames(names []string) ([]string, error) {

	ObjectIds := []string{}

	if mcs.getError != "" {
		if mcs.getError == "return_empty" {
			return ObjectIds, nil
		} else if mcs.getError == "not_found" {
			return ObjectIds, repository.Err_CategoryNotFound
		}

		return ObjectIds, nil
	}

	for _ = range names {
		ObjectIds = append(ObjectIds, utils.GenerateUUID())
	}
	return ObjectIds, nil

}

func TestBookService(t *testing.T) {

	// Test Save Book
	TestSaveBook(t)
	// Test GetBook
	// TestGetBook(t)

}

func TestSaveBook(t *testing.T) {
	TestSaveBookErrors(t)
	// We are checking the test book service here .
	mbr := repository.NewBookRepository(&MockBookRepository{nil})
	cs := NewCategoryServiceInterface(&MockCategoryService{"", ""})

	bs := NewBookService(*mbr, *cs)
	type inputStruct struct {
		name           string
		authors        []string
		co_authors     []string
		audiobook_urls []string
		ebook_urls     []string
		hard_copies    []string
		categories     []string
	}
	type test struct {
		want  error
		input inputStruct
	}
	table :=
		test{want: nil,
			input: inputStruct{
				name:           "david-and-goliath",
				authors:        []string{"david gog"},
				co_authors:     []string{},
				audiobook_urls: []string{},
				ebook_urls:     []string{},
				hard_copies:    []string{},
				categories:     []string{"comedy"}},
		}
	got := bs.SaveBook(table.input.name,
		table.input.authors, table.input.co_authors,
		table.input.audiobook_urls, table.input.ebook_urls, table.input.hard_copies, table.input.categories)
	if got != table.want {
		t.Errorf("Wanted %v but got %v", table.want, got)
	}

}
func TestSaveBookErrors(t *testing.T) {

	mbr := repository.NewBookRepository(&MockBookRepository{nil})
	cs := NewCategoryServiceInterface(&MockCategoryService{"", "return_empty"})
	cs1 := NewCategoryServiceInterface(&MockCategoryService{"", "return_empty"})
	cs2 := NewCategoryServiceInterface(&MockCategoryService{"", "not_found"})
	bs := NewBookService(*mbr, *cs)
	bs1 := NewBookService(*mbr, *cs1)
	bs2 := NewBookService(*mbr, *cs2)
	type inputStruct struct {
		name           string
		authors        []string
		co_authors     []string
		audiobook_urls []string
		ebook_urls     []string
		hard_copies    []string
		categories     []string
	}
	type test struct {
		want  error
		input inputStruct
		bs    *BookService
	}
	table := []test{
		{want: Err_Invalid_Category,
			input: inputStruct{
				name:           "david-and-goliath",
				authors:        []string{"david gog"},
				co_authors:     []string{},
				audiobook_urls: []string{},
				ebook_urls:     []string{},
				hard_copies:    []string{},
				categories:     []string{"comedy"}},
			bs: bs,
		},
		{want: Err_Invalid_Categories,
			input: inputStruct{
				name:           "david-and-goliath",
				authors:        []string{"david gog"},
				co_authors:     []string{},
				audiobook_urls: []string{},
				ebook_urls:     []string{},
				hard_copies:    []string{},
				categories:     []string{"comedy", "dark_comedy"}},
			bs: bs1,
		},
		{want: repository.Err_CategoryNotFound,
			input: inputStruct{
				name:           "david-and-goliath",
				authors:        []string{"david gog"},
				co_authors:     []string{},
				audiobook_urls: []string{},
				ebook_urls:     []string{},
				hard_copies:    []string{},
				categories:     []string{"comedy", "dark_comedy"}},
			bs: bs2,
		},
	}
	for _, object := range table {
		want := object.want
		got := object.bs.SaveBook(object.input.name,
			object.input.authors, object.input.co_authors,
			object.input.audiobook_urls, object.input.ebook_urls,
			object.input.hard_copies, object.input.categories)

		if got != want {
			t.Errorf("Wanted %v but got %v", want, got)
		}
	}
}
