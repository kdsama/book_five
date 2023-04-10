package mongodb

import (
	// "encoding/json"

	"fmt"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"go.mongodb.org/mongo-driver/bson"
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
