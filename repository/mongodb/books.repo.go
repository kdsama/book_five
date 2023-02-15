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
	"go.mongodb.org/mongo-driver/mongo"
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
			"uuid":        utils.GenerateUUID(),
			"name":        NewBook.Name,
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

func (g *MongoBookRepository) FindOrInsertBooksAndGetID(books []domain.Book) ([]string, []error, int) {

	errorSlice := []error{}
	idSlice := []string{}
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx := mongoUtils.GetQueryContext()
	errCount := 0
	for _, book := range books {
		var result domain.Book
		var err error
		if book.ID != "" {
			query := bson.M{"uuid": book.ID}

			err = col.FindOne(ctx, query).Decode(&result)
			if err != nil && err != mongo.ErrNoDocuments {
				errorSlice = append(errorSlice, err)
				errCount++
				idSlice = append(idSlice, "")

			}
		}

		uuid := utils.GenerateUUID()
		// If BookdId was not empty and we got NoDocument Error
		// This will barely happen .
		if err == mongo.ErrNoDocuments || (err == nil && book.ID == "") {
			_, err := col.InsertOne(
				ctx,
				bson.M{
					"uuid":        uuid,
					"name":        book.Name,
					"authors":     book.Authors,
					"co_authors":  book.Co_Authors,
					"audio_urls":  book.AudiobookUrls,
					"ebook_urls":  book.EbookUrls,
					"hard_copies": book.Hardcopies,
					"categories":  book.Categories,
					"createdAt":   book.Created_Timestamp,
					"updatedAt":   book.Updated_Timestamp,
					"verified":    false},
			)

			if err != nil {
				errorSlice = append(errorSlice, repository.ErrWriteRecord)
				idSlice = append(idSlice, "")
				errCount++

			} else {
				errorSlice = append(errorSlice, nil)
				idSlice = append(idSlice, uuid)
			}

		}
	}
	return idSlice, errorSlice, errCount
}
