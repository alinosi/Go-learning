package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type MockcommentRepositoryImpl struct {
	email, comment string
}

func NewCommentMockRepository(db *sql.DB) CommentRepository {
	return &MockcommentRepositoryImpl{email: "mockTesting@gmail.com", comment: "mockTesting"}
}

func (repository *MockcommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	
}

func (repository *MockcommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	
}

func (repository *MockcommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	
}
