package mongodb

import (
	// "encoding/json"

	"context"
	"errors"
	"time"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	// "errors"
	// "log"
	// "fmt"
)

var (
	Err_UserNotFound = errors.New("user couldnot be found")
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
		return ErrWriteRecord
	}
	return nil
}
