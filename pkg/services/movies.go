package services

import (
	"go-movies-crud/pkg/models"
	"go-movies-crud/pkg/repository"
)

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{
		repo: repo,
	}
}

func (s *MovieService) GetMovies() []models.Movie {
	return s.repo.GetMovies()
}

func (s *MovieService) GetMovieById(id string) (*models.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s *MovieService) CreateMovie(movie models.Movie) {
	s.repo.CreateMovie(movie)
}

func (s *MovieService) UpdateMovie(id string, updatedMovie models.Movie) error {
	return s.repo.UpdateMovie(id, updatedMovie)
}

func (s *MovieService) DeleteMovie(id string) error {
	return s.repo.DeleteMovie(id)
}
