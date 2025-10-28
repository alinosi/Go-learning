package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
)

type MockcommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentMockRepository(db *sql.DB) CommentRepository {
	return &MockcommentRepositoryImpl{DB: db}
}

func (repository *MockcommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	comment.Comment = "mock Insert berhasil dijalankan"
	return comment, nil
}

func (repository *MockcommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	comment := entity.Comment{Id: 0, Email: "mock@testing", Comment: "mock FindById berhasil dijalankan"}
	return comment, nil
}

func (repository *MockcommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	for i := 0; i < 3; i++ {
		comment := entity.Comment{Id: int32(i), Email: "mock@testing", Comment: "mock FindAll berhasil dijalankan"}
		comments = append(comments, comment)
	}

	return comments, nil
}
