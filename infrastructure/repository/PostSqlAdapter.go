package repository

import (
	"context"
	"example/core/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PostSqlAdapter struct {
	DB *gorm.DB
}

type Post struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	Deleted   bool
}

func NewPostSqlAdapter(db *gorm.DB) *PostSqlAdapter {
	if !db.Migrator().HasTable(Post{}) {
		err := db.Migrator().CreateTable(Post{}) //Is AutoMigrate better ??
		if err != nil {
			return &PostSqlAdapter{}
		}
	}
	return &PostSqlAdapter{DB: db}
}

func (a PostSqlAdapter) InsertPost(ctx context.Context, post model.Post) error {
	newPost := Post{
		CreatedAt: time.Now(),
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
	}

	result := a.DB.Create(&newPost)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a PostSqlAdapter) GetPosts(ctx context.Context) ([]model.Post, error) {
	models := []Post{}

	err := a.DB.Where("deleted = false").Find(&models).Error
	if err != nil {
		return []model.Post{}, err
	}

	domainModels := make([]model.Post, len(models))

	for i, post := range models {
		domainModels[i] = model.Post{ID: post.ID, Title: post.Title, Content: post.Content}
	}

	return domainModels, nil
}

func (a PostSqlAdapter) GetPost(ctx context.Context, postID uuid.UUID) (*model.Post, error) {
	post := Post{}

	err := a.DB.First(&post, "id = ? AND deleted = false", postID.String()).Error
	if err != nil {
		return &model.Post{}, err
	}
	return &model.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (a PostSqlAdapter) UpdatePost(ctx context.Context, post model.Post) error {
	newPost := Post{
		UpdatedAt: time.Now(),
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
	}

	err := a.DB.Updates(&newPost).Error
	if err != nil {
		return err
	}
	return nil
}

func (a PostSqlAdapter) DeletePost(ctx context.Context, postID uuid.UUID) error {
	//soft delete:
	newPost := Post{
		DeletedAt: time.Now(),
		Deleted:   true,
		ID:        postID,
	}
	err := a.DB.Updates(&newPost).Error
	if err != nil {
		return err
	}
	return nil
}
