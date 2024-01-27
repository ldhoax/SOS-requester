package main

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/sentry"
)

func main() {
	sentry.Init()
	server, err := api.NewServer()

	if err != nil {
		panic("Error when start server")
	}
	server.Run()
}
