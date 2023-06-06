package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kdsama/book_five/api"
	"github.com/kdsama/book_five/api/middleware"
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
	// categorySeeder(categoryservice)

	bookservice := service.NewBookService(*bookrepo, *categoryInterface)
	bookInterface := service.NewBookServiceInterface(bookservice)
	bookHandler := api.NewBookHandler(*bookInterface)

	usertokenMongoInstance := mongo_repo.NewMongoUserTokenRepository(mongoClient, "usertoken")
	usertokenrepo := repository.NewUserTokenRepository(usertokenMongoInstance)
	usertokenservice := service.NewUserTokenService(*usertokenrepo)
	usertokenserviceInterface := service.NewUserTokenServiceInterface(usertokenservice)

	userMongoInstance := mongo_repo.NewMongoUserRepository(mongoClient, "user")
	userrepo := repository.NewUserRepository(userMongoInstance)
	userservice := service.NewUserService(*userrepo, usertokenservice)
	userInterface := service.NewUserServiceInterface(userservice)
	userHandler := api.NewUserHandler(*userInterface)

	useractivityMongoInstance := mongo_repo.NewMongoUserActivityRepository(mongoClient, "useractivity")
	useractivityrepo := repository.NewUserActivityRepository(useractivityMongoInstance)
	useractivityservice := service.NewUserActivityService(userInterface, *useractivityrepo)
	userActivityInterface := service.NewUserActivityServiceInterface(useractivityservice)

	listcommentMongoInstance := mongo_repo.NewMongoListCommentRepository(mongoClient, "listcomments")
	listcommentrepo := repository.NewListCommentRepository(listcommentMongoInstance)
	listcommentservice := service.NewListCommentService(*listcommentrepo)
	listcommentserviceInterface := service.NewListCommentServiceInterface(listcommentservice)

	userlistMongoInstance := mongo_repo.NewMongoUserListRepository(mongoClient, "userlists")
	userlistrepo := repository.NewUserListRepository(userlistMongoInstance)

	userlistservice := service.NewUserListService(userInterface, bookInterface, userActivityInterface, listcommentserviceInterface, userlistrepo)
	userlistserviceInterface := service.NewUserListServiceInterface(userlistservice)
	userlistHandler := api.NewUserListHandler(*userlistserviceInterface)

	usertokenHandler := middleware.NewUserTokenHandler(*usertokenserviceInterface)

	// jobs.SaveBooks(bookInterface)
	router := http.NewServeMux()

	router.HandleFunc("/api/v1/book", bookHandler.Req)
	router.HandleFunc("/api/v1/user/login", userHandler.Req)
	router.HandleFunc("/api/v1/user/register", userHandler.Req)
	router.Handle("/api/v1/userlist", usertokenHandler.Authenticator(userlistHandler.Handler()))
	router.Handle("/api/v1/userlist/", usertokenHandler.Authenticator(userlistHandler.Handler()))

	log.Fatal(http.ListenAndServe(":8090", router))

}

func ConnectMongo(ctx context.Context) *mongodb.MongoClient {
	mongoClient := mongodb.GetMongoClient("mongodb://localhost:27017", "book_five")
	err := mongoClient.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}
