package main

import (
	"example/core/application/service"
	"example/infrastructure/repository"
	"example/presentation/rest"
)

func main() {
	// Adapters
	postRepository := repository.NewPostMemoryAdapter()

	// TODO: Implement repository.PostSqlAdapter and use it instead of repository.PostMemoryAdapter
	// postRepository := repository.NewPostSqlAdapter()

	// Services
	postService := service.NewPostService(postRepository)

	// Router
	router := rest.NewRouter(postService)

	router.Run()
}
