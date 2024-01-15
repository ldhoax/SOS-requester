package controllers

import (
	"github.com/sirupsen/logrus"
)

type Helper struct {
	logger *logrus.Logger
}

type Controller struct {
	Helper
}

func NewHelper(lg *logrus.Logger) *Helper {
	return &Helper{logger: lg}
}

func NewController(lg *logrus.Logger) Controller {
	helper := NewHelper(lg)
	return Controller{*helper}
}
