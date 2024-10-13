package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Controller) Health(c *gin.Context) {
	data := gin.H{
		"welcome": "Requester API alive",
	}
	s.Helper.Respond(c, "live", data, http.StatusOK)
}
