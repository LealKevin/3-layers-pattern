package category

import "fmt"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Store interface {
	GetAll() ([]Category, error)
	GetById(id int) (Category, error)
	Create(category Category) error
	Delete(id int) error
}

type MemoryStore struct {
	categories []Category
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		categories: []Category{
			{Id: 111, Name: "Food"},
			{Id: 222, Name: "Games"},
			{Id: 333, Name: "Cinema"},
		},
	}
}

func (s *MemoryStore) GetAll() ([]Category, error) {
	return s.categories, nil
}

func (s *MemoryStore) GetById(id int) (Category, error) {
	for _, categoty := range s.categories {
		if categoty.Id == id {
			return categoty, nil
		}
	}
	return Category{}, fmt.Errorf("category with ID %v not found", id)
}

func (s *MemoryStore) Create(category Category) error {
	s.categories = append(s.categories, category)
	return nil
}

func (s *MemoryStore) Delete(index int) error {
	s.categories = append(s.categories[:index], s.categories[index+1:]...)
	return nil
}
