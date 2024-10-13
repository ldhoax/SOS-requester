package service

import (
	"github.com/GoldenOwlAsia/go-golang-api/internal/request/model"
	"github.com/GoldenOwlAsia/go-golang-api/internal/request/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{repo: r}
}

func (s Service) GetAll() ([]model.Request, error) {
	return s.repo.FindAll()
}

func (s Service) Get(id string) (model.Request, error) {
	return s.repo.Find(id)
}
