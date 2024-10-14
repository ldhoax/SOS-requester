package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AuthController) Login(c *gin.Context) {
	type response struct {
		ID string `json:"id"`
		Username string `json:"username"`
		Token string `json:"token"`
	}

	data, err := a.authService.Login(c)

	if err != nil {
		a.helper.Respond(c, "Login fail", err, 403)
		return
	}

	a.helper.Respond(c, "Login success", response{
		ID: data.ID,
		Username: data.Username,
		Token: data.Token,
	}, http.StatusOK)
}
