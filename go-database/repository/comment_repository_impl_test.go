package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 13)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestMockCommentInsert(t *testing.T) {
	commentMockRepository := NewCommentMockRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "mockrepository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentMockRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestMockFindById(t *testing.T) {
	commentMockRepository := NewCommentMockRepository(belajar_golang_database.GetConnection())

	comment, err := commentMockRepository.FindById(context.Background(), 13)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestMockFindAll(t *testing.T) {
	commentMockRepository := NewCommentMockRepository(belajar_golang_database.GetConnection())

	comments, err := commentMockRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
