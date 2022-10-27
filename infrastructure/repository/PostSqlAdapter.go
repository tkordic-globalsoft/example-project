package repository

import (
	"context"
	"example/core/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostSqlAdapter struct {
	DB *gorm.DB
}

func NewPostSqlAdapter(db *gorm.DB) *PostSqlAdapter {
	return &PostSqlAdapter{DB: db}
}

func (a PostSqlAdapter) InsertPost(ctx context.Context, post model.Post) error {
	result := a.DB.Create(&post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a PostSqlAdapter) GetPosts(ctx context.Context) ([]model.Post, error) {
	models := []model.Post{}

	err := a.DB.Find(&models).Error
	if err != nil {
		return []model.Post{}, err
	}

	return models, nil
}

func (a PostSqlAdapter) GetPost(ctx context.Context, postID uuid.UUID) (*model.Post, error) {
	post := model.Post{}

	err := a.DB.First(&post, "id = ?", postID.String()).Error
	if err != nil {
		return &post, err
	}

	return &post, nil
}

func (a PostSqlAdapter) UpdatePost(ctx context.Context, post model.Post) error {
	err := a.DB.Save(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (a PostSqlAdapter) DeletePost(ctx context.Context, postID uuid.UUID) error {
	err := a.DB.Where("id = ?", postID.String()).Delete(&model.Post{}).Error
	if err != nil {
		return err
	}
	return nil
}
