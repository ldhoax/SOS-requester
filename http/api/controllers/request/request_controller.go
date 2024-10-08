package request

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api/controllers"
	RequestRepo "github.com/GoldenOwlAsia/go-golang-api/internal/request/repository"
	RequestService "github.com/GoldenOwlAsia/go-golang-api/internal/request/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RequestController struct {
	helper      *controllers.Helper
	requestService RequestService.Service
}

func NewRequestService(lg *logrus.Logger, model RequestService.Service) RequestController {
	helper := controllers.NewHelper(lg)
	return RequestController{
		helper:      helper,
		requestService: model,
	}
}

func NewRequestController(lg *logrus.Logger, db *gorm.DB) RequestController {
	return NewRequestService(lg, RequestService.NewService(RequestRepo.NewRepository(db)))
}
