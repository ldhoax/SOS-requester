package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AuthController) Login(c *gin.Context) {
	type response struct {
		Token string `json:"token"`
	}

	token, err := a.authService.Login(c)

	if err != nil {
		a.helper.Respond(c, "Login fail", err, 403)
		return
	}

	a.helper.Respond(c, "Login success", response{
		Token: token,
	}, http.StatusOK)
}
