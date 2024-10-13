package service

import (
	"github.com/ldhoax/SOS-requester/internal/user/model"
	"github.com/ldhoax/SOS-requester/pkg/erru"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CreateParams struct {
	Username string       `valid:"required"`
	Password string       `valid:"required"`
	Email    string       `valid:"required"`
	Status   model.Status `valid:"required"`
}

func (s Service) Create(ctx *gin.Context, params CreateParams) (model.User, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.User{}, erru.ErrArgument{Wrapped: err}
	}

	tx := s.repo.Db.Begin()
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.User{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
		Status:   params.Status,
	}
	err := s.repo.Create(&entity)
	if err != nil {
		return model.User{}, err
	}

	tx.Commit()
	return entity, err
}
