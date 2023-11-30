package main

import (
	"github.com/gin-gonic/gin"
	"go-movies-crud/pkg/handlers"
	"go-movies-crud/pkg/repository"
	"go-movies-crud/pkg/services"
)

func main() {
	router := gin.Default()

	movieInMemoryMovieRepository := repository.NewInMemoryMovieRepository()
	movieService := services.NewMovieService(movieInMemoryMovieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)

	router.GET("/movies", movieHandler.GetMovies)
	router.GET("/movies/:id", movieHandler.GetMovieById)
	router.POST("/movies", movieHandler.CreateMovie)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
