package service

import (
	"github.com/tulioguaraldo/go-microservice/core/movie/model"
	"github.com/tulioguaraldo/go-microservice/core/movie/repository"
)

type IMovieService interface {
	AllMovies() ([]model.Movie, error)
	FindMovieById(id uint) (*model.Movie, error)
	InsertMovie(movie *model.Movie) error
}

type service struct {
	repository repository.IMovieRepository
}

func NewMovieService(repository repository.IMovieRepository) IMovieService {
	return &service{
		repository: repository,
	}
}

func (s *service) AllMovies() ([]model.Movie, error) {
	return s.repository.All()
}

func (s *service) FindMovieById(id uint) (*model.Movie, error) {
	return s.repository.Find(id)
}

func (s *service) InsertMovie(movie *model.Movie) error {
	return s.repository.Insert(movie)
}
