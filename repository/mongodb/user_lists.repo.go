package mongodb

import (
	// "encoding/json"

	"context"
	"time"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
	"go.mongodb.org/mongo-driver/bson"
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
			"uuid":       utils.GenerateUUID(),
			"user_id":    user_list.User_ID,
			"book_ids":   user_list.Book_IDs,
			"name":       user_list.Name,
			"reaction":   user_list.Reactions,
			"comments":   user_list.Comments,
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
