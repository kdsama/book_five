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
	err_BookNotFound      = errors.New("Book couldnot be found ")
	err_NoBooksInCategory = errors.New("No Book is present in this category")
	errWriteRecord        = errors.New("cannot write to repository")
)

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
		bson.M{"name": NewBook.Name,
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
		return errWriteRecord
	}
	return nil
}

// func (g *MongoIpRepository) GetBook(Ip string) (*domain.IpDocument, error) {
// 	var result domain.IpDocument
// 	flag := false

// 	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
// 	ctx := mongoUtils.GetQueryContext()

// 	var r domain.IpDocument
// 	query := bson.M{"ip": Ip}
// 	err := col.FindOne(ctx, query).Decode(&r)
// 	if err != nil {

// 	} else {
// 		result = r
// 		var rr map[string]interface{}
// 		err := json.Unmarshal([]byte(result.SbrsData), &rr)
// 		if err != nil {
// 			continue
// 		}
// 		result.SbrsMapData = rr
// 		result.SbrsMapData["source"] = collection
// 		flag = true
// 		break

// 	}

// 	if !flag {
// 		return &result, errNoRecords
// 	}
// 	return &result, nil

// }
