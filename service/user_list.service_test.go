package service

import (
	"fmt"
	"testing"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
)

type mockUserListRepo struct {
	err error
}

var mockUserLists []domain.UserList

func (mul *mockUserListRepo) SaveUserList(item *domain.UserList) error {
	if mul.err != nil {
		return mul.err
	}
	mockUserLists = append(mockUserLists, *item)
	return nil
}

func (mul *mockUserListRepo) CountExistingListsOfAUser(user_id string) (int, error) {

	if mul.err != nil {
		return 0, mul.err
	}
	count := 0
	for i := range mockUserLists {
		if mockUserLists[i].User_ID == user_id {
			count++
		}
	}
	return count, nil
}

func Initialize() (*UserDI, MockBookRepository, *BookDI) {
	mc := MockCategoryService{}

	mci := NewCategoryServiceInterface(&mc)
	mci.SaveCategory("comedy", []string{})

	mrepo := &mockUserRepo{}
	usrepo := repository.NewUserRepository(mrepo)
	us := NewUserService(*usrepo)
	usi := NewUserServiceInterface(us)
	usi.SaveUser("kshitijdhingra@gmail.com", "kshitij", "e86f78a8a3caf0b60d8e74e5942aa6d86dc150cd3c03338aef25b7d2d7e3acc7")
	bookObject := MockBookRepository{}
	br := repository.NewBookRepository(&bookObject)

	bs := NewBookService(*br, *mci)
	bsi := NewBookServiceInterface(bs)
	return usi, bookObject, bsi
}

func TestSaveUserList(t *testing.T) {

	usi, bookObject, bsi := Initialize()
	mockrepo := &mockUserListRepo{}
	usls := NewUserListService(usi, bsi, NewUserActivityServiceInterface(&MockUserActivityService{}), mockrepo)

	SeedBooks(bookObject)

	// create a list of books now. book name 1,3,5,7,9 are going to be added here
	book_ids := []string{}
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			book_ids = append(book_ids, MockBooks[i].ID)
		}
	}

	id := "123"
	for i := range mockUsers {
		if mockUsers[i].Name == "kshitij" {

			id = mockUsers[i].ID
			break
		}
	}
	err := usls.SaveUserList(id, "This is a test summary / about ", "ListName-", book_ids)
	if err != nil {
		t.Errorf("Did not expect any error here but got %v", err)
	}
	r := mockUserLists[0].About

	if r != "This is a test summary / about " {
		t.Error("Expected matching abouts but not found ")
	}
	// usls.SaveUserList()

	// save 4 more times
	TestSaveUserListErrors(t)
	// Save another time there should be an erro r
}

func TestSaveUserListErrors(t *testing.T) {

	// initialize
	usi, bookObject, bsi := Initialize()
	mockrepo := &mockUserListRepo{}
	usls := NewUserListService(usi, bsi, NewUserActivityServiceInterface(&MockUserActivityService{}), mockrepo)

	// create a list
	SeedBooks(bookObject)
	id := "123"
	book_ids := []string{}
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			book_ids = append(book_ids, MockBooks[i].ID)
		}
	}
	// get id of user , required for saveUserList service function
	for i := range mockUsers {
		if mockUsers[i].Name == "kshitij" {

			id = mockUsers[i].ID
			break
		}
	}
	// trying to save list 6 times. Should not be possible
	for i := 0; i < 5; i++ {
		err := usls.SaveUserList(id, "This is a test summary / about ", "ListName-"+fmt.Sprintf("%d", i), book_ids)
		if err != nil {
			t.Errorf("Did not expect any error here but got %v", err)
		}
	}
	got := usls.SaveUserList(id, "This is a test summary / about ", "ListName-", book_ids)
	want := err_ListCreationNotAllowed
	if got != want {
		t.Errorf("want %v but got %v", want, got)
	}
	// Now test for a user thats not found in the database
	got = usls.SaveUserList("wrongID", "This is a test summary / about ", "ListName-78", book_ids)
	want = repository.Err_UserNotFound
	if got != want {
		t.Errorf("want %v but got %v", want, got)
	}
}
