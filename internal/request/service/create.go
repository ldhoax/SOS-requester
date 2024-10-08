package service

import (
	"github.com/GoldenOwlAsia/go-golang-api/internal/request/model"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/erru"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"fmt"
)

type CreateParams struct {
	PhoneNumber    string    `valid:"optional"`
	Email          string    `valid:"optional"`
	Location       string    `valid:"required"`
	ShortDescription string    `valid:"required"`
	EmergencyLevel int       `valid:"optional"`
	Latitude       float64   `valid:"optional"`
	Longitude      float64   `valid:"optional"`
	Description    string    `valid:"optional"`
}

func (s Service) Create(ctx *gin.Context, params CreateParams) (model.Request, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.Request{}, erru.ErrArgument{Wrapped: err}
	}

	if params.PhoneNumber == "" && params.Email == "" {
		return model.Request{}, erru.ErrArgument{Wrapped: fmt.Errorf("at least one of PhoneNumber or Email must be provided")}
	}

	tx := s.repo.Db.Begin()
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Request{
		PhoneNumber:    params.PhoneNumber,
		Email:          params.Email,
		EmergencyLevel: params.EmergencyLevel,
		Latitude:       params.Latitude,
		Longitude:      params.Longitude,
		Location:       params.Location,
		ShortDescription: params.ShortDescription,
		Description:    params.Description,
	}
	err := s.repo.Create(&entity)
	if err != nil {
		return model.Request{}, err
	}

	tx.Commit()
	return entity, err
}
