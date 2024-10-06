package router

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api/controllers"
	"github.com/GoldenOwlAsia/go-golang-api/http/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Route interface {
	Handler() gin.HandlerFunc
}

func Register(r *gin.Engine, lg *logrus.Logger, db *gorm.DB) {
	controller := controllers.NewController(lg)

	r.Use(middlewares.CORSMiddleware())
	r.Use(controller.MiddlewareLogger())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", controller.Health)
		NewAuthRouter(v1, lg, db)

		v1.Use(middlewares.JwtAuthMiddleware())
		{
			NewUserRouter(v1, lg, db)
		}
	}
}
