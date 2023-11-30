package services

import (
	"go-movies-crud/pkg/models"
	"go-movies-crud/pkg/repository"
)

type MovieService interface {
	GetMovies() []models.Movie
	GetMovieById(id string) (*models.Movie, error)
	CreateMovie(movie models.Movie)
	UpdateMovie(id string, updatedMovie models.Movie) error
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

func (s movieService) GetMovies() []models.Movie {
	return s.repo.GetMovies()
}

func (s movieService) GetMovieById(id string) (*models.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s movieService) CreateMovie(movie models.Movie) {
	s.repo.CreateMovie(movie)
}

func (s movieService) UpdateMovie(id string, updatedMovie models.Movie) error {
	return s.repo.UpdateMovie(id, updatedMovie)
}

func (s movieService) DeleteMovie(id string) error {
	return s.repo.DeleteMovie(id)
}
