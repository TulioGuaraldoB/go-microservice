package movie

import "github.com/gin-gonic/gin"

func GetRoutes() *gin.Engine {
	movieController := MovieDependency()

	router := gin.Default()

	movie := router.Group("movie")
	{
		movie.GET("", movieController.GetAllMovies)
		movie.GET(":id", movieController.GetMovieById)
		movie.POST("", movieController.CreateMovie)
	}

	return router
}
