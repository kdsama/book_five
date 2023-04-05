package mongodb

import (
	// "encoding/json"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "errors"
	// "log"
	// "fmt"
)

type MongoUserTokenRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoUserTokenRepository(m *mongoUtils.MongoClient, current string) *MongoUserTokenRepository {
	return &MongoUserTokenRepository{*m, current}
}

func (g *MongoUserTokenRepository) SaveUserToken(NewUserToken *domain.UserToken) error {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"user_id":    NewUserToken.User_ID,
			"token":      NewUserToken.Token,
			"created_at": NewUserToken.CreatedAt,
			"updated_at": NewUserToken.UpdatedAt,
		},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}

func (g *MongoUserTokenRepository) GetUserTokenByID(id string) (*domain.UserToken, error) {
	var result domain.UserToken

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	query := bson.M{"user_id": id}
	err := col.FindOne(ctx, query).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, repository.Err_UserTokenNotFound
		}
		return &result, err
	}
	return &result, nil
}
func (g *MongoUserTokenRepository) GetUserByToken(token string) (*domain.UserToken, error) {
	var result domain.UserToken

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	query := bson.M{"token": token}

	err := col.FindOne(ctx, query).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, repository.Err_UserTokenNotFound
		}
		return &result, err
	}
	return &result, nil
}
