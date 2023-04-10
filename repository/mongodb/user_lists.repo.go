package mongodb

import (
	// "encoding/json"

	"context"
	"time"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "errors"
	// "log"
	// "fmt"
)

type MongoUserListRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoUserListRepository(m *mongoUtils.MongoClient, current string) *MongoUserListRepository {
	return &MongoUserListRepository{*m, current}
}

func (g *MongoUserListRepository) SaveUserList(user_list *domain.UserList) error {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"uuid":       user_list.ID,
			"user_id":    user_list.User_ID,
			"book_ids":   user_list.Book_IDs,
			"name":       user_list.Name,
			"reaction":   user_list.Reactions,
			"created_at": user_list.CreatedAt,
			"updated_at": user_list.UpdatedAt,
		},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}

func (g *MongoUserListRepository) CountExistingListsOfAUser(user_id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	filter := bson.M{"uuid": user_id}
	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (g *MongoUserListRepository) GetListByID(list_id string) (*domain.UserList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	var result domain.UserList
	query := bson.M{"uuid": list_id}
	err := col.FindOne(ctx, query).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, repository.Err_UserListNotFound
		}
		return &result, err
	}
	return &result, nil
}
