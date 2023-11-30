package di

import (
	"go-movies-crud/pkg/handlers"
	"go-movies-crud/pkg/repository"
	"go-movies-crud/pkg/services"
)

func InitializeMovieDependencies() handlers.MovieHandler {
	movieInMemoryMovieRepository := repository.NewInMemoryMovieRepository()
	movieService := services.NewMovieService(movieInMemoryMovieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)

	return movieHandler
}
