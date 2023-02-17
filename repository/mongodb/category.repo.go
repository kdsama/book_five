package mongodb

import (
	// "encoding/json"

	"context"
	"log"
	"time"

	"github.com/kdsama/book_five/domain"
	mongoUtils "github.com/kdsama/book_five/infrastructure/mongodb"
	"github.com/kdsama/book_five/repository"
	"github.com/kdsama/book_five/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "errors"
	// "log"
	// "fmt"
)

type MongoCategoryRepository struct {
	repo    mongoUtils.MongoClient
	current string
}

func NewMongoCategoryRepository(m *mongoUtils.MongoClient, current string) *MongoCategoryRepository {
	return &MongoCategoryRepository{*m, current}
}

func (g *MongoCategoryRepository) SaveCategory(NewCategory *domain.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)

	_, err := col.InsertOne(
		ctx,
		bson.M{
			"uuid":           utils.GenerateUUID(),
			"name":           NewCategory.Name,
			"sub_categories": NewCategory.SubCategories,
			"createdAt":      NewCategory.Created_Timestamp,
			"updatedAt":      NewCategory.Updated_Timestamp},
	)

	if err != nil {
		return repository.ErrWriteRecord
	}
	return nil
}

func (g *MongoCategoryRepository) GetCategoryByID(id string) (*domain.Category, error) {
	var result domain.Category
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	query := bson.M{"uuid": id}
	err := col.FindOne(ctx, query).Decode(&result)
	if err != nil {
		log.Println(err)
		return &result, repository.Err_CategoryNotFound
	}

	return &result, nil

}
func (g *MongoCategoryRepository) GetCategoriesByManyIDs(ids []string) ([]domain.Category, error) {
	var result []domain.Category
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	query := bson.M{"uuid": bson.M{"$in": ids}}
	err := col.FindOne(ctx, query).Decode(&result)
	if err != nil {
		log.Println(err)
		return result, repository.Err_CategoryNotFound
	}

	return result, nil

}

func (g *MongoCategoryRepository) GetCategoriesByNames(names []string) ([]domain.Category, error) {
	var result []domain.Category
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	query := bson.M{"uuid": bson.M{"$in": names}}
	err := col.FindOne(ctx, query).Decode(&result)
	if err != nil {
		log.Println(err)
		return result, repository.Err_CategoryNotFound
	}

	return result, nil

}

func (g *MongoCategoryRepository) GetIdsByNames(names []string) ([]string, error) {

	var to_return []string
	col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	var results []domain.Category
	ctx, cancel := mongoUtils.GetQueryContext()
	defer cancel()

	filter := bson.M{"name": bson.M{"$in": names}}
	opts := options.Find().SetProjection(bson.M{"uuid": 1})
	cursor, err := col.Find(ctx, filter, opts)

	if err != nil {
		log.Println(err)
		return to_return, repository.Err_CategoryNotFound
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
		return to_return, repository.Err_CategoryNotFound
	}
	for _, category := range results {
		to_return = append(to_return, category.Id)
	}
	return to_return, nil

}
