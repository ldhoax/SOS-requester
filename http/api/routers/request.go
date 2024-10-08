package router

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api/controllers/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewRequestRouter(r *gin.RouterGroup, lg *logrus.Logger, db *gorm.DB) {
	requestController := request.NewRequestController(lg, db)

	request := r.Group("/requests")
	{
		request.GET("", requestController.Index)
		request.POST("", requestController.Create)
		request.GET("/:id", requestController.Get)
		// request.PUT("/:id", requestController.Update) // only allow update status
	}
}
