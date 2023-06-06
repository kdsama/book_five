package service

import (
	"errors"
	"fmt"

	"github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/entity"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
)

var (
	MAX_LIST_COUNT = int64(5)
	MAX_LIST_SIZE  = int64(5)
)

var (
	err_ListCannotBeCreated    = errors.New("there was an issue while creating the list")
	err_ListCreationNotAllowed = errors.New("Error list creation is not allowed")
	err_ListSizeExceeded       = errors.New("List size exceeds the maximum size")
	err_CannotComment          = errors.New("user is not allowed to comment on this list.")
	err_CannotReact            = errors.New("user is not allowed to react on the list")
)

type UserListService struct {
	book          BookServiceInterface
	user          UserServiceInterface
	user_activity UserActivityServiceInterface
	comment       ListCommentServiceInterface
	userlistRepo  repository.UserListRepo
}

func NewUserListService(user UserServiceInterface, book BookServiceInterface, user_activity UserActivityServiceInterface, list_comment ListCommentServiceInterface, userlistRepo repository.UserListRepo) *UserListService {

	return &UserListService{book, user, user_activity, list_comment, userlistRepo}
}

func (uls *UserListService) SaveUserList(user_id string, about string, list_name string, book_ids []string) error {

	user, err := uls.user.GetUserByID(user_id)
	if err != nil {
		// User just might not be present
		// this check probably will be done on middleware as well , the jwt would be checked here
		return err
	}
	if len(book_ids) > int(MAX_LIST_SIZE) {
		return err_ListSizeExceeded
	}

	//remove duplicates
	var book_mapping map[string]int
	new_book_ids := []string{}
	for i := range book_ids {
		if _, ok := book_mapping[book_ids[i]]; !ok {
			new_book_ids = append(new_book_ids, book_ids[i])
		}
	}
	// No Comment as it is a new List
	// No need to check if the book exist
	// We will save the book separately first and only then pass it to the user list
	// If they dont we need to create the books probably
	// Make them unverified. The books needs to be verified
	// Make a parameter , if that parameter is true, The user can create more than 5 lists,
	// lets say 20 lists. That parameter reveals for the particular user how many books
	// can he or she or they can add.
	// In the future it will be related to some kind of book score for that particular user

	timestamp := utils.GetCurrentTimestamp()

	userListObject := domain.NewUserList(user.ID, about, new_book_ids, list_name, timestamp)
	list_count, err := uls.CountExistingListsOfAUser(user_id)
	if err != nil {
		return err
	}
	// if one more added , we have to check whether the limit is reached or exceeded .
	if list_count+1 >= MAX_LIST_COUNT {
		return err_ListCreationNotAllowed
	}
	err = uls.userlistRepo.SaveUserList(userListObject)
	// no need to create an activity for creating the user list ???
	// for now lets skip it .
	return err
}

func (uls *UserListService) CountExistingListsOfAUser(user_id string) (int64, error) {

	count, err := uls.userlistRepo.CountExistingListsOfAUser(user_id)
	if err != nil && err != repository.Err_NoUserListFound {
		return 0, err
	}
	return count, nil

}

func (uls *UserListService) SaveComment(list_id string, user_id string, comment string) (string, error) {

	// check if list exists
	list, err := uls.userlistRepo.GetListByID(list_id)
	if err != nil {
		return "", err
	}

	// false check if user is allowed to do put a comment
	canComment := true

	if !canComment {
		return "", err_CannotComment
	}

	id, err := uls.comment.SaveListComment(list_id, user_id, comment)
	if err != nil {
		return "", err
	}
	err = uls.user_activity.SaveUserActivity(user_id, "comment", "", list.User_ID, list_id, "", "")
	if err != nil {
		return "", err
	}
	return id, nil
	//

}

func (uls *UserListService) GetComments(list_id string) ([]domain.ListComment, error) {
	// check if list exists
	_, err := uls.userlistRepo.GetListByID(list_id)
	if err != nil {
		return []domain.ListComment{}, err
	}

	return uls.comment.GetCommentsByListID(list_id)
}

func (uls *UserListService) React(user_id string, list_id string, reaction string, comment_id string) error {

	// Check if user-list exists,if it does return it
	list, err := uls.userlistRepo.GetListByID(list_id)
	if err != nil {
		return err
	}
	canReact := true
	if !canReact {
		return err_CannotReact
	}
	existing_reactions := list.Reactions
	previous_reaction, err := uls.user_activity.GetUserReactionActivityByUserAndListID(user_id, list_id)
	fmt.Printf("%+v %s \n ", previous_reaction, previous_reaction.Reaction)
	if err != nil && err != repository.Err_ActivityNotFound {
		return err
	}

	switch reaction {
	case "like":
		existing_reactions.Like++
	case "dislike":
		existing_reactions.DisLike++
	case "love":
		existing_reactions.Love++
	case "angry":
		existing_reactions.Angry++
	}
	switch previous_reaction.Reaction {
	case "like":
		existing_reactions.Like--
	case "dislike":
		existing_reactions.DisLike--
	case "love":
		existing_reactions.Love--
	case "angry":
		existing_reactions.Angry--
	}
	err_new := uls.UpdateUserListReactions(list_id, existing_reactions)
	if err_new != nil {
		return err
	}
	if err == repository.Err_ActivityNotFound {
		return uls.user_activity.SaveUserActivity(user_id, "reaction", reaction, list.User_ID, list_id, "", "")
	}

	return uls.user_activity.UpdateUserActivty(user_id, "reaction", reaction, list.User_ID, list_id, "", "")

	// check if the person already reacted
	// if the have already reacted on this
	// check which reaction it was
	// if its the same reaction dont do anything
	// if its not , undo the previous reaction , do the next reaction
	// save the updates on userlist
	// save user activity

}

func (uls *UserListService) UpdateUserListReactions(list_id string, reaction entity.Reaction) error {
	return uls.userlistRepo.UpdateUserListReactions(list_id, reaction)
}
