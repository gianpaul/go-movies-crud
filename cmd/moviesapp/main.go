package main

import (
	"github.com/gin-gonic/gin"
	"go-movies-crud/pkg/di"
	"go-movies-crud/pkg/routers"
)

func main() {
	router := gin.Default()

	movieHandler := di.InitializeMovieDependencies()

	routers.NewMoviesRouter(movieHandler, router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
