package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/GoldenOwlAsia/go-golang-api/configs"
	"gopkg.in/gomail.v2"
)

type Data struct {
	Info interface{}
}

type Mailer struct {
	subject  string
	template string
	from     string
	to       string
	username string
	password string
	port     int
	host     string
}

func NewMailer(c configs.Config) *Mailer {
	var mailer Mailer
	mailer.from = c.Mailer.FromAddress
	mailer.username = c.Mailer.UserName
	mailer.password = c.Mailer.Password
	mailer.host = c.Mailer.Host
	mailer.port = c.Mailer.Port
	return &mailer
}

func (m Mailer) SendMailTemplate(i Data) error {
	var err error
	if len(m.template) == 0 {
		err = errors.New("template required")
		return err
	}

	// t := template.New(m.template)

	t, err := template.ParseFiles(m.template)
	if err != nil {
		return (err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i.Info); err != nil {
		return (err)
	}
	result := tpl.String()

	msg := gomail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(m.from, "No-reply"))
	msg.SetHeader("To", m.to)

	msg.SetHeader("Subject", m.subject)
	msg.SetBody("text/html", result)

	d := gomail.NewDialer(m.host, m.port, m.username, m.password)

	if err := d.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func (m Mailer) SendMailMessage(message string) error {
	fmt.Print(m.from)
	msg := gomail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(m.from, "No-reply"))
	msg.SetHeader("To", m.to)
	msg.SetHeader("Subject", m.subject)
	msg.SetBody("text/html", message)

	d := gomail.NewDialer(m.host, m.port, m.username, m.password)
	if err := d.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func (m *Mailer) SetTemplate(template string) (*Mailer, error) {
	path, _ := os.Getwd()
	html := filepath.Join(path, template)
	_, err := os.Stat(html)
	if err != nil {
		return m, err
	}

	m.template = html
	return m, nil
}

func (m *Mailer) SetSubject(subject string) *Mailer {
	m.subject = subject
	return m
}

func (m *Mailer) To(mail string) *Mailer {
	m.to = mail
	return m
}
