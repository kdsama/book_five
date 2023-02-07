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

type MongoUserRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoUserRepository(m *mongoUtils.MongoClient, current string) *MongoUserRepository {
	return &MongoUserRepository{*m, current}
}

func (g *MongoUserRepository) SaveUser(NewUser *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"email":      NewUser.Email,
			"created_at": NewUser.CreatedAt,
			"updated_at": NewUser.UpdatedAt,
			"pwd":        NewUser.Password,
		},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}

func (g *MongoUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var result domain.User

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx := mongoUtils.GetQueryContext()

	query := bson.M{"email": email}
	err := col.FindOne(ctx, query).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, repository.Err_UserNotFound
		}
		return &result, err
	}
	return &result, nil
}
