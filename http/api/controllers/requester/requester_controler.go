package request

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api/controllers"
	RequesterRepo "github.com/GoldenOwlAsia/go-golang-api/internal/requester/repository"
	RequesterService "github.com/GoldenOwlAsia/go-golang-api/internal/requester/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RequestController struct {
	helper      *controllers.Helper
	requesterService RequesterService.Service
}

func NewRequestService(lg *logrus.Logger, model RequestService.Service) RequestController {
	helper := controllers.NewHelper(lg)
	return RequestController{
		helper:      helper,
		requesterService: model,
	}
}

func NewRequesterController(lg *logrus.Logger, db *gorm.DB) RequesterController {
	return NewRequesterService(lg, RequesterService.NewService(RequesterRepo.NewRepository(db)))
}
