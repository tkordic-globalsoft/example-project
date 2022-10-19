package repository

import (
	"context"
	"example/core/domain/model"
	"fmt"
	"github.com/google/uuid"
)

type PostMemoryAdapter struct {
	data map[string]model.Post
}

func NewPostMemoryAdapter() *PostMemoryAdapter {
	data := make(map[string]model.Post)

	return &PostMemoryAdapter{
		data: data,
	}
}

func (a PostMemoryAdapter) InsertPost(ctx context.Context, post model.Post) error {
	key := post.ID.String()
	a.data[key] = post

	return nil
}

func (a PostMemoryAdapter) GetPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post

	for _, value := range a.data {
		posts = append(posts, value)
	}

	return posts, nil
}

func (a PostMemoryAdapter) GetPost(ctx context.Context, postID uuid.UUID) (*model.Post, error) {
	key := postID.String()

	post, ok := a.data[key]
	if !ok {
		return &model.Post{}, fmt.Errorf("unable to find post with ID '%s'", postID)
	}

	return &post, nil
}

func (a PostMemoryAdapter) UpdatePost(ctx context.Context, post model.Post) error {
	key := post.ID.String()
	a.data[key] = post

	return nil
}

func (a PostMemoryAdapter) DeletePost(ctx context.Context, postID uuid.UUID) error {
	key := postID.String()
	delete(a.data, key)

	return nil
}
