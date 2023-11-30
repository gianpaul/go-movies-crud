package router

import (
	"github.com/gin-gonic/gin"
	"go-movies-crud/feature/movies/handler"
)

func NewMoviesRouter(movieHandler handler.MovieHandler, router *gin.Engine) {
	movies := router.Group("/movies")
	{
		movies.GET("/", movieHandler.GetMovies)
		movies.GET("/:id", movieHandler.GetMovieById)
		movies.POST("/", movieHandler.CreateMovie)
	}
}
