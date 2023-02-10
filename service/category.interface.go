package service

type CategoryServiceInterface interface {
	SaveCategory(name string, categories []string) error
	GetIdsByNames(names []string) ([]string, error)
}

type CategoryDI struct {
	CategoryServiceInterface
}

func NewCategoryServiceInterface(br CategoryServiceInterface) *CategoryDI {
	return &CategoryDI{br}
}
