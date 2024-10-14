package service

import (
	"errors"

	"github.com/ldhoax/SOS-requester/internal/requester/model"
	"github.com/ldhoax/SOS-requester/internal/utils/token"
	"github.com/ldhoax/SOS-requester/pkg/erru"
	"github.com/ldhoax/SOS-requester/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type ResponseData struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Token string `json:"token"`
}

func (s Service) Login(c *gin.Context) (ResponseData, error) {
	var loginInput LoginData

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		return ResponseData{}, erru.ErrArgument{Wrapped: err}
	}

	if _, err := govalidator.ValidateStruct(loginInput); err != nil {
		return ResponseData{}, erru.ErrArgument{Wrapped: err}
	}

	u := model.Requester{}
	u.Username = loginInput.Username
	u.Password = loginInput.Password

	response, err := s.LoginCheck(&u)
	if err != nil {
		return ResponseData{}, err
	}

	return response, nil
}

func (s Service) LoginCheck(entity *model.Requester) (ResponseData, error) {
	var err error

	u := model.Requester{}

	err = s.repo.Db.Model(model.Requester{}).Where("username = ?", entity.Username).Take(&u).Error

	if err != nil {
		return ResponseData{}, err
	}

	err = utils.VerifyPassword(entity.Password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return ResponseData{}, errors.New("Username or password incorrect")
	} else {
		if err != nil {
			return ResponseData{}, err
		}
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return ResponseData{}, err
	}

	return ResponseData{
		ID: u.ID,
		Username: u.Username,
		Token: token,
	}, nil
}
