package service

import (
	"testing"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

type MockUserTokenService struct{ err error }

var mockTokens []domain.UserToken

func (muts *MockUserTokenService) GenerateUserToken(user_id string) (string, error) { return "", nil }
func (muts *MockUserTokenService) GenerateAndSaveUserToken(user_id string) (string, error) {
	token, err := utils.CreateJWTToken(user_id)
	if err != nil {
		return "", err
	}
	mockTokens = append(mockTokens, domain.UserToken{User_ID: user_id, Token: token})
	return token, nil
}
func (muts *MockUserTokenService) GetUserTokenByID(user_id string) (*domain.UserToken, error) {
	if muts.err != nil {
		return &domain.UserToken{}, muts.err
	}
	for i := range mockTokens {
		if mockTokens[i].User_ID == user_id {
			return &mockTokens[i], nil
		}
	}
	return &domain.UserToken{}, repository.Err_UserTokenNotFound
}
func (muts *MockUserTokenService) ValidateUserTokenAndGetUserID(token string) (string, error) {
	return "", nil
}

// -------------------------------------------------------------------------------
type MockUserTokenRepo struct{ err error }

func (mutr *MockUserTokenRepo) SaveUserToken(token *domain.UserToken) error {
	if mutr.err != nil {
		return mutr.err
	}
	mockTokens = append(mockTokens, *token)
	return nil
}

func (mutr *MockUserTokenRepo) GetUserTokenByID(user_id string) (*domain.UserToken, error) {
	if mutr.err != nil {
		return &domain.UserToken{}, mutr.err
	}
	for i := range mockTokens {
		if mockTokens[i].User_ID == user_id {
			return &mockTokens[i], nil
		}
	}
	return &domain.UserToken{}, repository.Err_UserTokenNotFound
}

func (mutr *MockUserTokenRepo) GetUserByToken(token string) (*domain.UserToken, error) {
	if mutr.err != nil {
		return &domain.UserToken{}, mutr.err
	}
	for i := range mockTokens {
		if mockTokens[i].Token == token {
			return &mockTokens[i], nil
		}
	}
	return &domain.UserToken{}, repository.Err_UserTokenNotFound
}

// -----------------------------------------------------------------

func TestSaveUserToken(t *testing.T) {
	repo := repository.NewUserTokenRepository(&MockUserTokenRepo{})
	usertokenservice := NewUserTokenService(*repo)
	mockTokens = []domain.UserToken{}
	user_id := "kshitij@gmail.com"
	token, err := usertokenservice.GenerateAndSaveUserToken(user_id)
	if err != nil {
		t.Errorf("did not expect error, but got %v", err)
	}
	// check if token is present or not for the correct user
	flag := false
	for i := range mockTokens {
		if mockTokens[i].Token == token {
			if mockTokens[i].User_ID == user_id {
				flag = true
			}

		}
	}
	if !flag {
		t.Error("cannot find token ")
	}

}

func TestValidateUserTokenAndGetUserIDErrors(t *testing.T) {
	repo1 := repository.NewUserTokenRepository(&MockUserTokenRepo{})
	usertokenservice1 := NewUserTokenService(*repo1)
	mockTokens = []domain.UserToken{}

	correctToken, err := usertokenservice1.GenerateAndSaveUserToken("kshitij@gmail.com")
	if err != nil {
		t.Errorf("did not expect error, but got %v", err)
	}

	// change timestamp of this person

	mockTokens[0].UpdatedAt -= 37000
	_, got := usertokenservice1.ValidateUserTokenAndGetUserID(correctToken)
	want := Err_TokenExpired
	if want != got {
		t.Errorf("2-->wanted %v but got %v", want, got)
	}

	// change user_id of this person
	mockTokens[0].User_ID = "someoneElse@gmail.com"
	_, got = usertokenservice1.ValidateUserTokenAndGetUserID(correctToken)
	want = Err_MaliciousIntent
	if want != got {
		t.Errorf("2-->wanted %v but got %v", want, got)
	}

	repo := repository.NewUserTokenRepository(&MockUserTokenRepo{repository.Err_UserTokenNotFound})
	usertokenservice := NewUserTokenService(*repo)
	//usertokennot found ,
	dummytoken, err := utils.CreateJWTToken("something@gmail.com")
	if err != nil {
		t.Errorf("did not expect error, but got %v", err)
	}

	_, got = usertokenservice.ValidateUserTokenAndGetUserID(dummytoken)
	want = repository.Err_UserTokenNotFound
	if want != got {
		t.Errorf("1-->wanted %v but got %v", want, got)
	}
	// Jwt error, string passes is not a valid jwt token
	_, got = usertokenservice.ValidateUserTokenAndGetUserID("thisisnotavalidtokenyoushouldknowthat")
	want = utils.Err_InvalidToken
	if want != got {
		t.Errorf("2-->wanted %v but got %v", want, got)
	}

}
