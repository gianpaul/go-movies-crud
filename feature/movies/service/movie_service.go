package service

import (
	"go-movies-crud/feature/movies/model"
	"go-movies-crud/feature/movies/repository"
)

type MovieService interface {
	GetMovies() []model.Movie
	GetMovieById(id string) (*model.Movie, error)
	CreateMovie(movie model.Movie)
	UpdateMovie(id string, updatedMovie model.Movie) error
	DeleteMovie(id string) error
}

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

func (s movieService) GetMovies() []model.Movie {
	return s.repo.GetMovies()
}

func (s movieService) GetMovieById(id string) (*model.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s movieService) CreateMovie(movie model.Movie) {
	s.repo.CreateMovie(movie)
}

func (s movieService) UpdateMovie(id string, updatedMovie model.Movie) error {
	return s.repo.UpdateMovie(id, updatedMovie)
}

func (s movieService) DeleteMovie(id string) error {
	return s.repo.DeleteMovie(id)
}
