package repository

import (
	"github.com/tulioguaraldo/go-microservice/core/movie/model"
	"gorm.io/gorm"
)

type IMovieRepository interface {
	All() ([]model.Movie, error)
	Find(id uint) (*model.Movie, error)
	Insert(movie *model.Movie) error
}

type repository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) IMovieRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) All() ([]model.Movie, error) {
	movies := []model.Movie{}
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *repository) Find(id uint) (*model.Movie, error) {
	movie := model.Movie{}
	if err := r.db.First(&movie, &id).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *repository) Insert(movie *model.Movie) error {
	return r.db.Create(movie).Error
}
