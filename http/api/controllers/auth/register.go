package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AuthController) Register(c *gin.Context) {
	_, err := a.authService.Register(c)
	if err != nil {
		a.helper.Respond(c, "Register fail", err, 403)
		return
	}

	a.helper.Respond(c, "Register success", nil, http.StatusOK)
}
