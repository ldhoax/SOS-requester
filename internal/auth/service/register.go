package service

import (
	"github.com/GoldenOwlAsia/go-golang-api/configs"
	"github.com/GoldenOwlAsia/go-golang-api/internal/user/model"
	UserService "github.com/GoldenOwlAsia/go-golang-api/internal/user/service"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/erru"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/mailer"
	"github.com/gin-gonic/gin"
)

func (s Service) Register(c *gin.Context) (uint, error) {
	var registerInput UserService.CreateParams
	userService := UserService.NewService(s.repo)
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		return 0, erru.ErrArgument{Wrapped: err}
	}
	// need to verify email to active account
	registerInput.Status = 2
	newUser, err := userService.Create(c, registerInput)
	if err != nil {
		return 0, err
	}
	// recommendation use goroutine to send mail queue
	_ = s.SendMailVerify(newUser)
	return newUser.ID, nil
}

func (s Service) SendMailVerify(u model.User) error {
	var err error

	cnf, err := configs.NewParsedConfig()
	if err != nil {
		return err
	}
	// create service linkActive and paste this
	var dataMail = map[string]string{
		"username":   u.Username,
		"linkActive": "linkk..............",
	}
	m := mailer.NewMailer(cnf)

	m, err = m.To(u.Email).SetTemplate("web/views/mails/register.html")
	if err != nil {
		return err
	}
	err = m.SetSubject("Active your account").SendMailTemplate(mailer.Data{Info: dataMail})

	return err

}
