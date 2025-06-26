package category

import "fmt"

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetAll() ([]Category, error) {
	return s.Store.GetAll()
}

func (s *Service) GetById(id int) (Category, error) {
	category, err := s.Store.GetById(id)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

func (s *Service) Create(category Category) error {
	characterLimit := 3
	if len(category.Name) < characterLimit {
		return fmt.Errorf("category name can't have less then: %v characters ", characterLimit)
	}
	err := s.Store.Create(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id int, categories []Category) error {
	for i, category := range categories {
		if category.Id == id {
			s.Store.Delete(i)
			return nil
		}
	}
	return fmt.Errorf("unable to find category with id: %v", id)
}
