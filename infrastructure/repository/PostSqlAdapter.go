package repository

import (
	"context"
	"example/core/domain/model"
	"github.com/google/uuid"
)

type PostSqlAdapter struct {
}

func NewPostSqlAdapter() *PostSqlAdapter {
	return &PostSqlAdapter{}
}

func (a PostSqlAdapter) InsertPost(ctx context.Context, post model.Post) error {
	// TODO: Implement
	return nil
}

func (a PostSqlAdapter) GetPosts(ctx context.Context) ([]model.Post, error) {
	// TODO: Implement
	return []model.Post{}, nil
}

func (a PostSqlAdapter) GetPost(ctx context.Context, postID uuid.UUID) (*model.Post, error) {
	// TODO: Implement
	return &model.Post{}, nil
}

func (a PostSqlAdapter) UpdatePost(ctx context.Context, post model.Post) error {
	// TODO: Implement
	return nil
}

func (a PostSqlAdapter) DeletePost(ctx context.Context, postID uuid.UUID) error {
	// TODO: Implement
	return nil
}
