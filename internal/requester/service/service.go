package service

import (
	"github.com/GoldenOwlAsia/go-golang-api/internal/requester/model"
	"github.com/GoldenOwlAsia/go-golang-api/internal/requester/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{repo: r}
}

func (s Service) Get(id string) (model.Requester, error) {
	return s.repo.Find(id)
}

func (s Service) Create(entity *model.Requester) error {
	return s.repo.Create(entity)
}