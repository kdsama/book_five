package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	bookservice := service.NewBookService(*bookrepo, *categoryservice)

	categoryservice.SaveCategory("drama", []string{})
	categoryservice.SaveCategory("action", []string{})
	categoryservice.SaveCategory("comedy", []string{})
	categoryservice.SaveCategory("fiction", []string{})
	categoryservice.SaveCategory("thriller", []string{})

	bookservice.SaveBook("I'd Like to Play Alone, Please: Essays", []string{"Tom Segura"},
		[]string{},
		[]string{"https://www.amazon.co.uk/Id-Like-Play-Alone-Please/dp/B09QXP8GD1/ref=tmm_aud_swatch_0?_encoding=UTF8&qid=&sr="},
		[]string{}, []string{"https://www.amazon.co.uk/Id-Like-Play-Alone-Please/dp/1538704633"}, []string{"comedy"})
	fmt.Println("So this is done ")
	os.Exit(0)

}
func ConnectMongo(ctx context.Context) *mongodb.MongoClient {
	mongoClient := mongodb.GetMongoClient("mongodb://localhost:27017", "book_five")
	err := mongoClient.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}
