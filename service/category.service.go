package service

import (
	"time"

	domain "github.com/kdsama/book_five/domain"
	"github.com/kdsama/book_five/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(category repository.CategoryRepository) *CategoryService {
	return &CategoryService{category}
}
func (bs *CategoryService) SaveCategory(name string, categories []string) error {
	// validation checks already done in the http handler .
	// We just need to save it for now.
	timestamp := time.Now().Unix()
	// Get all Ids of categories
	CategoryObjects := []primitive.ObjectID{}
	var err error
	if len(categories) != 0 {
		CategoryObjects, err = bs.categoryRepo.GetIdsByNames(categories)
		if err != nil {
			//
			return err
		}
	}
	CategoryObject := domain.NewCategory(name, CategoryObjects, timestamp)
	bs.categoryRepo.SaveCategory(CategoryObject)
	return nil

}

func (bs *CategoryService) GetIdsByNames(names []string) ([]primitive.ObjectID, error) {
	return bs.categoryRepo.GetIdsByNames(names)
}
