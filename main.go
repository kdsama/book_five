package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kdsama/book_five/api"
	"github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	mongo_repo "github.com/kdsama/book_five/repository/mongodb"
	"github.com/kdsama/book_five/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := ConnectMongo(ctx)
	bookMongoInstance := mongo_repo.NewMongoBookRepository(mongoClient, "books")
	bookrepo := repository.NewBookRepository(bookMongoInstance)

	categoryMongoInstance := mongo_repo.NewMongoCategoryRepository(mongoClient, "categories")
	categoryrepo := repository.NewCategoryRepository(categoryMongoInstance)
	categoryservice := service.NewCategoryService(*categoryrepo)
	categoryInterface := service.NewCategoryServiceInterface(categoryservice)
	bookservice := service.NewBookService(*bookrepo, *categoryInterface)
	bookInterface := service.NewBookServiceInterface(bookservice)
	// categorySeeder(categoryservice)

	bookHandler := api.NewBookHandler(*bookInterface)
	http.HandleFunc("/api/v1/book", bookHandler.Req)

	log.Fatal(http.ListenAndServe(":8090", nil))

}

func ConnectMongo(ctx context.Context) *mongodb.MongoClient {
	mongoClient := mongodb.GetMongoClient("mongodb://localhost:27017", "book_five")
	err := mongoClient.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}
