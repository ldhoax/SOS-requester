package router

import (
	"github.com/ldhoax/SOS-requester/http/api/controllers/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewUserRouter(r *gin.RouterGroup, lg *logrus.Logger, db *gorm.DB) {
	userController := user.NewUserController(lg, db)
	user := r.Group("/user")
	{
		user.GET("", userController.CurrentUser)
		user.GET("/:id", userController.Get)
		user.POST("", userController.Create)

	}
}
