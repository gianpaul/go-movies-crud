package main

import (
	"github.com/gin-gonic/gin"
	"go-movies-crud/feature/movies/di"
	"go-movies-crud/feature/movies/router"
)

func main() {
	routers := gin.Default()

	movieHandler := di.InitializeMovieDependencies()

	router.NewMoviesRouter(movieHandler, routers)

	err := routers.Run("localhost:8080")
	if err != nil {
		return
	}
}
