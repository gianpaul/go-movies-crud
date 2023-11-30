package repository

import (
	"errors"
	"go-movies-crud/pkg/models"
)

type MovieRepository interface {
	GetMovies() []models.Movie
	GetMovieById(id string) (*models.Movie, error)
	CreateMovie(movie models.Movie)
	UpdateMovie(id string, updatedMovie models.Movie) error
	DeleteMovie(id string) error
}

type InMemoryMovieRepository struct {
	movies []models.Movie
}

func NewInMemoryMovieRepository() *InMemoryMovieRepository {
	return &InMemoryMovieRepository{
		movies: []models.Movie{},
	}
}

func (r *InMemoryMovieRepository) GetMovies() []models.Movie {
	return r.movies
}

func (r *InMemoryMovieRepository) GetMovieById(id string) (*models.Movie, error) {
	for _, movie := range r.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

func (r *InMemoryMovieRepository) CreateMovie(movie models.Movie) {
	r.movies = append(r.movies, movie)
}

func (r *InMemoryMovieRepository) UpdateMovie(id string, updatedMovie models.Movie) error {
	for i, movie := range r.movies {
		if movie.ID == id {
			r.movies[i] = updatedMovie
			return nil
		}
	}
	return errors.New("movie not found")
}

func (r *InMemoryMovieRepository) DeleteMovie(id string) error {
	for i, movie := range r.movies {
		if movie.ID == id {
			r.movies = append(r.movies[:i], r.movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}
