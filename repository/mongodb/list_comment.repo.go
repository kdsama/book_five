package mongodb

import (
	// "encoding/json"

	"context"
	"fmt"
	"log"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "errors"
	// "log"
	// "fmt"
)

type MongoListCommentRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoListCommentRepository(m *mongoUtils.MongoClient, current string) *MongoListCommentRepository {
	return &MongoListCommentRepository{*m, current}
}

func (g *MongoListCommentRepository) SaveListComment(list_comment *domain.ListComment) error {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()
	fmt.Println("are we here or not ???")
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"uuid":       list_comment.ID,
			"user_id":    list_comment.User_ID,
			"list_id":    list_comment.List_ID,
			"comment":    list_comment.Comment,
			"reaction":   list_comment.Reaction,
			"created_at": list_comment.CreatedAt,
			"updated_at": list_comment.UpdatedAt,
		},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}

func (g *MongoListCommentRepository) GetCommentsByListID(list_id string) ([]domain.ListComment, error) {
	return g.GetCommentsByListIDOpts(list_id, 20, 0)
}

func (g *MongoListCommentRepository) GetCommentsByListIDOpts(list_id string, size, skip int) ([]domain.ListComment, error) {
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	var results []domain.ListComment

	filter := bson.M{"list_id": list_id}
	opts := options.Find().SetProjection(bson.M{"list_id": 1})
	opts.SetLimit(int64(size))
	opts.SetSkip(int64(skip))
	cursor, err := col.Find(ctx, filter, opts)

	if err != nil {
		log.Println(err)
		return results, repository.Err_CannotLoadComments
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
		return results, repository.Err_CannotLoadComments
	}
	return results, nil

}
