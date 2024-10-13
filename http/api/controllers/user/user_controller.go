package user

import (
	"github.com/ldhoax/SOS-requester/http/api/controllers"
	UserRepo "github.com/ldhoax/SOS-requester/internal/user/repository"
	UserService "github.com/ldhoax/SOS-requester/internal/user/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserController struct {
	helper      *controllers.Helper
	userService UserService.Service
}

func NewUserService(lg *logrus.Logger, model UserService.Service) UserController {
	helper := controllers.NewHelper(lg)
	return UserController{
		helper:      helper,
		userService: model,
	}
}

func NewUserController(lg *logrus.Logger, db *gorm.DB) UserController {
	return NewUserService(lg, UserService.NewService(UserRepo.NewRepository(db)))
}
