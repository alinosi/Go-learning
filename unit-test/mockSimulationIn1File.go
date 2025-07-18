package main

// this is a simulation for mock impelemenation but in one file

import (
	"errors"
)

// Category-entity
type Category struct {
	Id   string
	Name string
}

// CategoryRepository
type CategoryRepository interface {
	FindById(id string) *Category
}

// CategoryService
type CategoryService struct {
	Repository CategoryRepository // bisa berisi apapun yang penting impelemen CategoryRepository
}

func (service CategoryService) Get(id string) (*Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}

func main() {
	test := CategoryService{}

	test.Get("1")
}

// this is contain code from service, entity, and repository directory
