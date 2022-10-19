package rest

import (
	"example/core/application/service"
	"example/core/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (c PostController) IndexPosts(ctx *gin.Context) {
	posts, err := c.postService.IndexPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response []postDTO
	for _, post := range posts {
		response = append(response, postDTOFromModel(post))
	}

	ctx.JSON(http.StatusOK, response)
}

func (c PostController) StorePost(ctx *gin.Context) {
	var dto storePostRequestDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	post, err := c.postService.StorePost(ctx, dto.Title, dto.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := postDTOFromModel(*post)

	ctx.JSON(http.StatusCreated, response)
}

func (c PostController) ShowPost(ctx *gin.Context) {
	postID := ctx.Param("ID")

	post, err := c.postService.ShowPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := postDTOFromModel(*post)

	ctx.JSON(http.StatusOK, response)
}

func (c PostController) UpdatePost(ctx *gin.Context) {
	postID := ctx.Param("ID")

	var dto updatePostRequestDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	post, err := c.postService.UpdatePost(ctx, postID, dto.Title, dto.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := postDTOFromModel(*post)

	ctx.JSON(http.StatusOK, response)
}

func (c PostController) DestroyPost(ctx *gin.Context) {
	postID := ctx.Param("ID")

	if err := c.postService.DestroyPost(ctx, postID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

type storePostRequestDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type updatePostRequestDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type postDTO struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

func postDTOFromModel(post model.Post) postDTO {
	return postDTO{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}
