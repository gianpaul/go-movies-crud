package handler

import (
	"github.com/gin-gonic/gin"
	"go-movies-crud/feature/movies/model"
	"go-movies-crud/feature/movies/service"
	"net/http"
)

type MovieHandler interface {
	GetMovies(c *gin.Context)
	GetMovieById(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
}

type movieHandler struct {
	service service.MovieService
}

func NewMovieHandler(service service.MovieService) MovieHandler {
	return &movieHandler{
		service: service,
	}
}

func (h *movieHandler) GetMovies(c *gin.Context) {
	movies := h.service.GetMovies()
	c.JSON(http.StatusOK, movies)
}

func (h *movieHandler) GetMovieById(c *gin.Context) {
	id := c.Param("id")
	movie, err := h.service.GetMovieById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (h *movieHandler) CreateMovie(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.CreateMovie(movie)
	c.JSON(http.StatusCreated, movie)
}

func (h *movieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateMovie(id, movie)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (h *movieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "movie deleted"})
}
