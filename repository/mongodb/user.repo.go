package mongodb

import (
	// "encoding/json"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
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
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"uuid":       utils.GenerateUUID(),
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

func (g *MongoUserRepository) GetUserById(id string) (*domain.User, error) {
	var result domain.User

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	query := bson.M{"uuid": id}
	err := col.FindOne(ctx, query).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, repository.Err_UserNotFound
		}
		return &result, err
	}
	return &result, nil
}

func (g *MongoUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var result domain.User

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

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

func (g *MongoBookRepository) CountUsersFromIDs(user_ids []string) (int64, error) {

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	filter := bson.M{"uuid": bson.M{"$in": user_ids}}
	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, err
}
