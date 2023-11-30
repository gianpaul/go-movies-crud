package routers

import (
	"github.com/gin-gonic/gin"
	"go-movies-crud/pkg/handlers"
)

func NewMoviesRouter(movieHandler handlers.MovieHandler, router *gin.Engine) {
	movies := router.Group("/movies")
	{
		movies.GET("/", movieHandler.GetMovies)
		movies.GET("/:id", movieHandler.GetMovieById)
		movies.POST("/", movieHandler.CreateMovie)
	}
}
