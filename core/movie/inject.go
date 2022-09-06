package movie

import (
	"github.com/tulioguaraldo/go-microservice/app/db"
	"github.com/tulioguaraldo/go-microservice/core/movie/repository"
	"github.com/tulioguaraldo/go-microservice/core/movie/service"
)

func MovieDependency() controller {
	db := db.OpenConnection()
	repository := repository.NewMovieRepository(db)
	service := service.NewMovieService(repository)
	return NewMovieController(service)
}
