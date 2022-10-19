package service

import (
	"context"
	"example/core/domain/model"
	"example/core/port/repository"
	"fmt"
	"github.com/google/uuid"
)

type PostService struct {
	postRepository repository.PostPort
}

func NewPostService(postRepository repository.PostPort) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}

func (s PostService) IndexPosts(ctx context.Context) ([]model.Post, error) {
	posts, err := s.postRepository.GetPosts(ctx)
	if err != nil {
		return []model.Post{}, fmt.Errorf("postRepository.GetPosts error: %w", err)
	}

	return posts, nil
}

func (s PostService) StorePost(ctx context.Context, title string, content string) (*model.Post, error) {
	postUUID := uuid.New()

	post := model.Post{
		ID:      postUUID,
		Title:   title,
		Content: content,
	}

	if err := s.postRepository.InsertPost(ctx, post); err != nil {
		return &model.Post{}, fmt.Errorf("postRepository.InsertPost error: %w", err)
	}

	storedPost, err := s.postRepository.GetPost(ctx, postUUID)
	if err != nil {
		return &model.Post{}, fmt.Errorf("postRepository.GetPost error: %w", err)
	}

	return storedPost, nil
}

func (s PostService) ShowPost(ctx context.Context, postID string) (*model.Post, error) {
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return &model.Post{}, fmt.Errorf("uuid.Parse error: %w", err)
	}

	post, err := s.postRepository.GetPost(ctx, postUUID)
	if err != nil {
		return &model.Post{}, fmt.Errorf("postRepository.GetPost error: %w", err)
	}

	return post, nil
}

func (s PostService) UpdatePost(ctx context.Context, postID string, title string, content string) (*model.Post, error) {
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return &model.Post{}, fmt.Errorf("uuid.Parse error: %w", err)
	}

	post := model.Post{
		ID:      postUUID,
		Title:   title,
		Content: content,
	}

	if err := s.postRepository.UpdatePost(ctx, post); err != nil {
		return &model.Post{}, fmt.Errorf("postRepository.UpdatePost error: %w", err)
	}

	updatedPost, err := s.postRepository.GetPost(ctx, postUUID)
	if err != nil {
		return &model.Post{}, fmt.Errorf("postRepository.GetPost error: %w", err)
	}

	return updatedPost, nil
}

func (s PostService) DestroyPost(ctx context.Context, postID string) error {
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return fmt.Errorf("uuid.Parse error: %w", err)
	}

	if err = s.postRepository.DeletePost(ctx, postUUID); err != nil {
		return fmt.Errorf("postRepository.DeletePost error: %w", err)
	}

	return nil
}
