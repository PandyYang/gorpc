package service

import (
	"fmt"
	"gorpc/iris/repository"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repository.MovieRepository
}

func NewMovieServiceManager(repo repository.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	return fmt.Sprintln(m.repo.GetMovieName())
}
