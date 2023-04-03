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

var ()

type MongoBookRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoBookRepository(m *mongoUtils.MongoClient, current string) *MongoBookRepository {
	return &MongoBookRepository{*m, current}
}

func (g *MongoBookRepository) SaveBook(NewBook *domain.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"uuid":        NewBook.ID,
			"name":        NewBook.Name,
			"image_url":   NewBook.Image_Url,
			"authors":     NewBook.Authors,
			"co_authors":  NewBook.Co_Authors,
			"audio_urls":  NewBook.AudiobookUrls,
			"ebook_urls":  NewBook.EbookUrls,
			"hard_copies": NewBook.Hardcopies,
			"categories":  NewBook.Categories,
			"createdAt":   NewBook.Created_Timestamp,
			"updatedAt":   NewBook.Updated_Timestamp,
			"verified":    false},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}
