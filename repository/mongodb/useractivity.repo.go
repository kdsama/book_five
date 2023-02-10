package mongodb

import (
	// "encoding/json"

	"context"
	"time"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"go.mongodb.org/mongo-driver/bson"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"user_id":    user_activity.User_ID,
			"action":     user_activity.Action,
			"receiver":   user_activity.Receiver,
			"comment_id": user_activity.Comment_ID,
			"list_id":    user_activity.List_ID,
			"created_at": user_activity.CreatedAt,
			"updated_at": user_activity.UpdatedAt,
		},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}
