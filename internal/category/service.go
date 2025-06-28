package category

import "fmt"

type Service struct{}

func NewService(store Store) *Service {
	return &Service{}
}

func (s *Service) Create(name string) (Category, error) {
	characterLimit := 3
	if len(name) < characterLimit {
		return Category{}, fmt.Errorf("category name can't have less then: %v characters ", characterLimit)
	}
	var category Category
	category.Id = 1
	category.Name = name

	return category, nil
}
