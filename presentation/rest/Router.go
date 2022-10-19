package rest

import (
	"example/core/application/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	ginRouter      *gin.Engine
	postController *PostController
}

func NewRouter(postService *service.PostService) *Router {
	postController := NewPostController(postService)

	router := &Router{
		postController: postController,
	}

	defaultGinRouter := gin.Default()
	router.setupRoutes(defaultGinRouter)
	router.ginRouter = defaultGinRouter

	return router
}

func (r Router) setupRoutes(ginRouter *gin.Engine) {
	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "example service is running...",
		})
	})

	postGroup := ginRouter.Group("/posts")
	{
		postGroup.GET("", r.postController.IndexPosts)
		postGroup.POST("", r.postController.StorePost)
		postGroup.GET(":ID", r.postController.ShowPost)
		postGroup.PATCH(":ID", r.postController.UpdatePost)
		postGroup.DELETE(":ID", r.postController.DestroyPost)
	}
}

func (r Router) Run() {
	if err := r.ginRouter.Run(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}
