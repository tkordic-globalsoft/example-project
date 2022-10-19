package repository

import (
	"context"
	"example/core/domain/model"
	"github.com/google/uuid"
)

type PostPort interface {
	InsertPost(ctx context.Context, post model.Post) error              // Create
	GetPosts(ctx context.Context) ([]model.Post, error)                 // Read
	GetPost(ctx context.Context, postID uuid.UUID) (*model.Post, error) // Read
	UpdatePost(ctx context.Context, post model.Post) error              // Update
	DeletePost(ctx context.Context, postID uuid.UUID) error             // Delete
}
