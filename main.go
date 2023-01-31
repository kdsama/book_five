package main

import (
	"context"
	"log"
	"time"

	"github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	mongo_repo "github.com/kdsama/book_five/repository/mongodb"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := ConnectMongo(ctx)
	bookMongoInstance := mongo_repo.NewMongoBookRepository(mongoClient, "books")
	bookrepo := repository.NewBookRepository(bookMongoInstance)

}
func ConnectMongo(ctx context.Context) *mongodb.MongoClient {
	mongoClient := mongodb.GetMongoClient("mongodb://localhost:27017", "book_five")
	err := mongoClient.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}
