package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Controller) Health(c *gin.Context) {
	data := gin.H{
		"welcome": "Welcome to GoldenOwn Consulting",
	}
	s.Helper.Respond(c, "live", data, http.StatusOK)
}
