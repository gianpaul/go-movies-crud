package di

import (
	"go-movies-crud/feature/movies/handler"
	"go-movies-crud/feature/movies/repository"
	"go-movies-crud/feature/movies/service"
)

func InitializeMovieDependencies() handler.MovieHandler {
	movieInMemoryMovieRepository := repository.NewInMemoryMovieRepository()
	movieService := service.NewMovieService(movieInMemoryMovieRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	return movieHandler
}
