package main

import (
	"errors"
	"fmt"
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
	Repository CategoryRepository
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
	fmt.Println("hello wolrd")
}
