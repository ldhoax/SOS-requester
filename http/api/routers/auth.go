package router

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api/controllers/auth"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewAuthRouter(r *gin.RouterGroup, lg *logrus.Logger, db *gorm.DB) {
	authController := auth.NewAuthController(lg, db)
	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}
}
