package repository

import (
	"errors"
	"go-movies-crud/feature/movies/model"
)

type MovieRepository interface {
	GetMovies() []model.Movie
	GetMovieById(id string) (*model.Movie, error)
	CreateMovie(movie model.Movie)
	UpdateMovie(id string, updatedMovie model.Movie) error
	DeleteMovie(id string) error
}

type InMemoryMovieRepository struct {
	movies []model.Movie
}

func NewInMemoryMovieRepository() *InMemoryMovieRepository {
	return &InMemoryMovieRepository{
		movies: []model.Movie{},
	}
}

func (r *InMemoryMovieRepository) GetMovies() []model.Movie {
	return r.movies
}

func (r *InMemoryMovieRepository) GetMovieById(id string) (*model.Movie, error) {
	for _, movie := range r.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

func (r *InMemoryMovieRepository) CreateMovie(movie model.Movie) {
	r.movies = append(r.movies, movie)
}

func (r *InMemoryMovieRepository) UpdateMovie(id string, updatedMovie model.Movie) error {
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
