package api

import (
	"fmt"

	"github.com/GoldenOwlAsia/go-golang-api/configs"
	router "github.com/GoldenOwlAsia/go-golang-api/http/api/routers"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/db"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/sentry"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	logger   *logrus.Logger
	database *gorm.DB
	config   configs.Config
}

func NewServer() (*Server, error) {
	cnf, err := configs.NewParsedConfig()
	if err != nil {
		return nil, err
	}

	database, err := db.Connect(db.ConfingDB{
		Host:     cnf.Database.Host,
		Port:     cnf.Database.Port,
		User:     cnf.Database.User,
		Password: cnf.Database.Password,
		Name:     cnf.Database.Name,
	})
	if err != nil {
		sentry.Log(err)
		return nil, err
	}

	log := NewLogger()

	s := Server{
		logger:   log,
		config:   cnf,
		database: database,
	}
	return &s, nil
}

func (s *Server) GetRouter() *gin.Engine {
	r := gin.Default()
	router.Register(r, s.logger, s.database)
	r.Use(sentrygin.New(sentrygin.Options{}))
	return r
}

func (s *Server) Run() {
	// gin.SetMode(gin.ReleaseMode)
	serverPort := s.config.ServerPort
	r := s.GetRouter()
	err := r.Run(fmt.Sprintf("localhost:%d", serverPort))

	if err != nil {
		sentry.Log(err)
		return
	}
}
