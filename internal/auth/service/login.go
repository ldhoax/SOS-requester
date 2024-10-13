package service

import (
	"errors"

	"github.com/GoldenOwlAsia/go-golang-api/internal/requester/model"
	"github.com/GoldenOwlAsia/go-golang-api/internal/utils/token"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/erru"
	"github.com/GoldenOwlAsia/go-golang-api/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func (s Service) Login(c *gin.Context) (string, error) {
	var loginInput LoginData

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		return "", erru.ErrArgument{Wrapped: err}
	}

	if _, err := govalidator.ValidateStruct(loginInput); err != nil {
		return "", erru.ErrArgument{Wrapped: err}
	}

	u := model.Requester{}
	u.Username = loginInput.Username
	u.Password = loginInput.Password

	token, err := s.LoginCheck(&u)
	return token, err
}

func (s Service) LoginCheck(entity *model.Requester) (string, error) {
	var err error

	u := model.Requester{}

	err = s.repo.Db.Model(model.Requester{}).Where("username = ?", entity.Username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(entity.Password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("Username or password incorrect")
	} else {
		if err != nil {
			return "", err
		}
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
