package movie

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tulioguaraldo/go-microservice/core/movie/model"
	"github.com/tulioguaraldo/go-microservice/core/movie/service"
	"gorm.io/gorm"
)

type controller struct {
	service service.IMovieService
}

func NewMovieController(service service.IMovieService) controller {
	return controller{
		service: service,
	}
}

func (c *controller) GetAllMovies(ctx *gin.Context) {
	movies, err := c.service.AllMovies()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (c *controller) GetMovieById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	movie, err := c.service.FindMovieById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (c *controller) CreateMovie(ctx *gin.Context) {
	movieReq := model.Movie{}
	if err := ctx.ShouldBindJSON(&movieReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.service.InsertMovie(&movieReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "movie inserted successfully!",
		"movie":   movieReq,
	})
}
