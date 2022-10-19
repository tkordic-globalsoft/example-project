package service

import (
	"context"
	"example/core/domain/model"
	repositoryAdapter "example/infrastructure/repository"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func getPostService() *PostService {
	postRepository := repositoryAdapter.NewPostMemoryAdapter()

	ctx := context.Background()
	posts := getDummyPosts()
	for _, post := range posts {
		_ = postRepository.InsertPost(ctx, post)
	}

	return NewPostService(postRepository)
}

func TestPostService_IndexPosts(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		service *PostService
		args    args
		want    []model.Post
		wantErr bool
	}{
		{
			name:    "List all posts",
			service: getPostService(),
			args: args{
				ctx: context.Background(),
			},
			want:    getDummyPosts(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.service.IndexPosts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexPosts() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexPosts() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPostService_ShowPost(t *testing.T) {
	type args struct {
		ctx    context.Context
		postID string
	}

	tests := []struct {
		name    string
		service *PostService
		args    args
		want    *model.Post
		wantErr bool
	}{
		{
			name:    "Show post",
			service: getPostService(),
			args: args{
				ctx:    context.Background(),
				postID: "00000000-0000-0000-0000-000000000001",
			},
			want:    getDummyPost1(),
			wantErr: false,
		},
		{
			name:    "Show post that doesn't exist",
			service: getPostService(),
			args: args{
				ctx:    context.Background(),
				postID: "00000000-0000-0000-0000-000000000004",
			},
			want:    &model.Post{},
			wantErr: true,
		},
		{
			name:    "Show post with invalid ID",
			service: getPostService(),
			args: args{
				ctx:    context.Background(),
				postID: "00000000-0000-0000-0000-00000000000z",
			},
			want:    &model.Post{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.service.ShowPost(tt.args.ctx, tt.args.postID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowPost() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowPost() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPostService_StorePost(t *testing.T) {
	type args struct {
		ctx     context.Context
		title   string
		content string
	}

	tests := []struct {
		name    string
		service *PostService
		args    args
		want    *model.Post
		wantErr bool
	}{
		{
			name:    "Store post",
			service: getPostService(),
			args: args{
				ctx:     context.Background(),
				title:   "Post 4 Title",
				content: "Post 4 Content",
			},
			want:    getDummyPost4(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.service.StorePost(tt.args.ctx, tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorePost() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			tt.want.ID = got.ID

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorePost() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPostService_UpdatePost(t *testing.T) {
	type args struct {
		ctx     context.Context
		postID  string
		title   string
		content string
	}

	tests := []struct {
		name    string
		service *PostService
		args    args
		want    *model.Post
		wantErr bool
	}{
		{
			name:    "Update post",
			service: getPostService(),
			args: args{
				ctx:     context.Background(),
				postID:  "00000000-0000-0000-0000-000000000002",
				title:   "Updated Post 2 Title",
				content: "Updated Post 2 Content",
			},
			want:    getUpdatedDummyPost2(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.service.UpdatePost(tt.args.ctx, tt.args.postID, tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePost() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPostService_DestroyPost(t *testing.T) {
	type args struct {
		ctx    context.Context
		postID string
	}

	tests := []struct {
		name    string
		service *PostService
		args    args
		want    *model.Post
		wantErr bool
	}{
		{
			name:    "Destroy post",
			service: getPostService(),
			args: args{
				ctx:    context.Background(),
				postID: "00000000-0000-0000-0000-000000000001",
			},
			want:    &model.Post{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.service.DestroyPost(tt.args.ctx, tt.args.postID); (err != nil) != tt.wantErr {
				t.Errorf("DestroyPost() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			postUUID, _ := uuid.Parse(tt.args.postID)
			getPostWantErr := !tt.wantErr

			got, err := tt.service.postRepository.GetPost(tt.args.ctx, postUUID)
			if (err != nil) != getPostWantErr {
				t.Errorf("postRepository.GetPost() error = %+v, getPostWantErr %+v", err, getPostWantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postRepository.GetPost() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func getDummyPost1() *model.Post {
	ID, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")

	return &model.Post{
		ID:      ID,
		Title:   "Post 1 Title",
		Content: "Post 1 Content",
	}
}

func getDummyPost2() *model.Post {
	ID, _ := uuid.Parse("00000000-0000-0000-0000-000000000002")

	return &model.Post{
		ID:      ID,
		Title:   "Post 2 Title",
		Content: "Post 2 Content",
	}
}

func getUpdatedDummyPost2() *model.Post {
	ID, _ := uuid.Parse("00000000-0000-0000-0000-000000000002")

	return &model.Post{
		ID:      ID,
		Title:   "Updated Post 2 Title",
		Content: "Updated Post 2 Content",
	}
}

func getDummyPost3() *model.Post {
	ID, _ := uuid.Parse("00000000-0000-0000-0000-000000000003")

	return &model.Post{
		ID:      ID,
		Title:   "Post 3 Title",
		Content: "Post 3 Content",
	}
}

func getDummyPost4() *model.Post {
	ID, _ := uuid.Parse("00000000-0000-0000-0000-000000000004")

	return &model.Post{
		ID:      ID,
		Title:   "Post 4 Title",
		Content: "Post 4 Content",
	}
}

func getDummyPosts() []model.Post {
	return []model.Post{
		*getDummyPost1(),
		*getDummyPost2(),
		*getDummyPost3(),
	}
}
