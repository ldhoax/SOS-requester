package user

import (
	"net/http"

	userService "github.com/GoldenOwlAsia/go-golang-api/internal/user/service"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/erru"
	"github.com/gin-gonic/gin"
)

func (u UserController) Create(c *gin.Context) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type response struct {
		ID uint `json:"id"`
	}

	var userInput request
	if err := c.ShouldBindJSON(&userInput); err != nil {
		u.helper.Respond(c, "Create user fail", erru.ErrArgument{Wrapped: err}, 0)
	}
	user, err := u.userService.Create(c, userService.CreateParams{
		Username: userInput.Username,
		Password: userInput.Password,
		Email:    userInput.Email,
		Status:   1,
	})
	if err != nil {
		u.helper.Respond(c, "Create user fail", err, 0)
		return
	}
	u.helper.Respond(c, "Create user success", response{ID: user.ID}, http.StatusOK)

}
