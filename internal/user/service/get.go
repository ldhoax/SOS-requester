package service

import (
	"errors"

	"github.com/GoldenOwlAsia/go-golang-api/internal/user/model"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/db"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/erru"
	"github.com/gin-gonic/gin"
)

func (s Service) Get(ctx *gin.Context, id uint) (model.User, error) {
	user, err := s.repo.Find(id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.User{}, erru.ErrArgument{Wrapped: errors.New("user object not found")}
	default:
		return model.User{}, err
	}
	return user, nil
}

func (s Service) GetById(id uint) (model.User, error) {
	user, err := s.repo.Find(id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.User{}, erru.ErrArgument{Wrapped: errors.New("user object not found")}
	default:
		return model.User{}, err
	}
	prepairOutput(&user)
	return user, nil
}

func prepairOutput(u *model.User) {
	u.Password = ""
}
