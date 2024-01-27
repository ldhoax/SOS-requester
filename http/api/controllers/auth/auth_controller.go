package auth

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api/controllers"
	AuthService "github.com/GoldenOwlAsia/go-golang-api/internal/auth/service"
	UserRepo "github.com/GoldenOwlAsia/go-golang-api/internal/user/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthController struct {
	helper      *controllers.Helper
	authService AuthService.Service
}

func NewAuthService(lg *logrus.Logger, model AuthService.Service) AuthController {
	helper := controllers.NewHelper(lg)
	return AuthController{
		helper:      helper,
		authService: model,
	}
}

func NewAuthController(lg *logrus.Logger, db *gorm.DB) AuthController {
	return NewAuthService(lg, AuthService.NewService(UserRepo.NewRepository(db)))
}
