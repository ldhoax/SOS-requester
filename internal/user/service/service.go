package service

import (
	"github.com/ldhoax/SOS-requester/internal/user/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}
