package mongodb

import (
	// "encoding/json"

	"context"
	"fmt"
	"time"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "errors"
	// "log"
	// "fmt"
)

type MongoUserActivityRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoUserActivityRepository(m *mongoUtils.MongoClient, current string) *MongoUserActivityRepository {
	return &MongoUserActivityRepository{*m, current}
}

func (g *MongoUserActivityRepository) SaveUserActivity(user_activity *domain.UserActivity) error {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"uuid":       utils.GenerateUUID(),
			"user_id":    user_activity.User_ID,
			"action":     user_activity.Action,
			"reaction":   user_activity.Reaction,
			"receiver":   user_activity.Receiver,
			"comment_id": user_activity.Comment_ID,
			"list_id":    user_activity.List_ID,
			"created_at": user_activity.CreatedAt,
			"desc":       user_activity.Desc,
			"updated_at": user_activity.UpdatedAt,
		},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}

func (g *MongoUserActivityRepository) GetLastUserActivityByUserID(user_id string) (*domain.UserActivity, error) {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	query := bson.M{"user_id": user_id}
	var result domain.UserActivity
	queryOptions := options.FindOneOptions{}

	queryOptions.SetSort(bson.M{"user_id": -1})
	err := col.FindOne(ctx, query, &queryOptions).Decode(&result)

	return &result, err
}

func (g *MongoUserActivityRepository) GetUserReactionActivityByUserAndListID(user_id string, list_id string) (*domain.UserActivity, error) {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	query := bson.M{"user_id": user_id, "list_id": list_id, "action": "reaction"}
	var result domain.UserActivity
	queryOptions := options.FindOneOptions{}

	queryOptions.SetSort(bson.M{"user_id": -1})
	err := col.FindOne(ctx, query, &queryOptions).Decode(&result)

	if err == mongo.ErrNoDocuments {

		return &result, repository.Err_ActivityNotFound
	}
	return &result, err
}

func (g *MongoUserActivityRepository) UpdateUserActivty(dua *domain.UserActivity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	fmt.Println("???? how come we cannot find this ???", dua.User_ID, dua.List_ID, dua.Receiver)
	filter := bson.M{"user_id": dua.User_ID, "list_id": dua.List_ID, "reciever": dua.Receiver}

	update := bson.M{"$set": bson.M{"reaction": dua.Reaction, "action": dua.Action}}

	result, err := col.UpdateOne(ctx, filter, update)
	print(result)
	if err != nil {
		return err
	}
	return nil
}
