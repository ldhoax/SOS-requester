package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ldhoax/SOS-requester/internal/user/model"
	"github.com/ldhoax/SOS-requester/internal/utils/token"
	"github.com/ldhoax/SOS-requester/pkg/erru"
	"github.com/gin-gonic/gin"
)

func (u UserController) CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)
	fmt.Print(user_id)
	if err != nil {
		u.helper.Respond(c, err.Error(), nil, http.StatusBadRequest)
		return
	}

	getResponse, err := u.userService.GetById(uint(user_id))

	if err != nil {
		u.helper.Respond(c, err.Error(), nil, http.StatusBadRequest)
		return
	}

	u.helper.Respond(c, "Get user success", getResponse, http.StatusOK)
}

func (u UserController) Get(c *gin.Context) {
	type response struct {
		ID        uint         `json:"id"`
		Username  string       `json:"username"`
		Email     string       `json:"email"`
		Status    model.Status `json:"status"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at,omitempty"`
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		u.helper.Respond(c, "Get user fail", erru.ErrArgument{
			Wrapped: errors.New("valid id must provide in path"),
		}, 0)
		return
	}
	getResponse, err := u.userService.Get(c, uint(id))

	if err != nil {
		u.helper.Respond(c, "Get user fail", err, http.StatusNotFound)
		return
	}
	u.helper.Respond(c, "Get user sucess", response{
		ID:        getResponse.ID,
		Username:  getResponse.Username,
		Email:     getResponse.Email,
		Status:    getResponse.Status,
		CreatedAt: getResponse.CreatedAt,
		UpdatedAt: getResponse.UpdatedAt,
	}, http.StatusOK)

}
