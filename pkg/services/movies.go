package services

import (
	"errors"
	"go-movies-crud/pkg/models"
)

type MovieService struct {
	movies []models.Movie
}

func NewMovieService() *MovieService {
	return &MovieService{
		movies: []models.Movie{},
	}
}

func (s *MovieService) GetMovies() []models.Movie {
	return s.movies
}

func (s *MovieService) GetMovieById(id string) (*models.Movie, error) {
	for _, movie := range s.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

func (s *MovieService) CreateMovie(movie models.Movie) {
	s.movies = append(s.movies, movie)
}

func (s *MovieService) UpdateMovie(id string, updatedMovie models.Movie) error {
	for i, movie := range s.movies {
		if movie.ID == id {
			s.movies[i] = updatedMovie
			return nil
		}
	}
	return errors.New("movie not found")
}

func (s *MovieService) DeleteMovie(id string) error {
	for i, movie := range s.movies {
		if movie.ID == id {
			s.movies = append(s.movies[:i], s.movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}
