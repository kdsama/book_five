package service

import (
	"fmt"
	"testing"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/entity"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
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

func TestSaveUserList(t *testing.T) {
	// save a category
	mc := MockCategoryService{}

	mci := NewCategoryServiceInterface(&mc)
	mci.SaveCategory("comedy", []string{})
	// save a user
	mrepo := &mockUserRepo{}
	usrepo := repository.NewUserRepository(mrepo)
	us := NewUserService(*usrepo)
	usi := NewUserServiceInterface(us)
	usi.SaveUser("kshitijdhingra@gmail.com", "kshitij", "e86f78a8a3caf0b60d8e74e5942aa6d86dc150cd3c03338aef25b7d2d7e3acc7")
	type bookStruct struct {
		name           string
		authors        []string
		co_authors     []string
		audiobook_urls []string
		ebook_urls     []string
		hard_copies    []string
		categories     []string
	}

	book := &bookStruct{name: "david-and-goliath",
		authors:        []string{"david gog"},
		co_authors:     []string{},
		audiobook_urls: []string{},
		ebook_urls:     []string{},
		hard_copies:    []string{},
		categories:     []string{"comedy"}}
	bookObject := MockBookRepository{}

	for i := 0; i < 10; i++ {
		book.name = book.name + fmt.Sprintf("%d", i)
		obj := domain.NewBook(book.name, book.authors, book.co_authors,
			entity.MakeUrlObjects(book.audiobook_urls), entity.MakeUrlObjects(book.ebook_urls), entity.MakeUrlObjects(book.hard_copies), book.categories, utils.GetCurrentTimestamp())
		bookObject.SaveBook(obj)
	}

	// create a list of books now. book name 1,3,5,7,9 are going to be added here
	book_ids := []string{}
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			book_ids = append(book_ids, MockBooks[i].ID)
		}
	}
	br := repository.NewBookRepository(&bookObject)
	// userlist service
	bs := NewBookService(*br, *mci)
	bsi := NewBookServiceInterface(bs)
	usls := NewUserListService(usi, bsi, NewUserActivityServiceInterface(&MockUserActivityService{}), &mockUserListRepo{})
	id := "123"
	for i := range mockUsers {
		if mockUsers[i].Name == "kshitij" {

			id = mockUsers[i].ID

		}
	}
	err := usls.SaveUserList(id, "something something meri jaan ", "first list", book_ids)
	if err != nil {
		t.Errorf("Did not expect any error here but got %v", err)
	}
	r := mockUserLists[0].About

	if r != "something something meri jaan " {
		t.Error("Expected matching abouts but not found ")
	}
	// usls.SaveUserList()

	// save 4 more times
	for i := 0; i < 4; i++ {
		err := usls.SaveUserList(id, "something something meri jaan ", "first list"+fmt.Sprintf("%d", i), book_ids)
		if err != nil {
			t.Errorf("Did not expect any error here but got %v", err)
		}
	}
	err = usls.SaveUserList(id, "something something meri jaan ", "first list", book_ids)
	if err != err_ListCreationNotAllowed {
		t.Errorf("want %v but got %v", err_ListCreationNotAllowed, err)
	}
	// Save another time there should be an erro r
}
