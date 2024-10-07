package main

import (
	"github.com/GoldenOwlAsia/go-golang-api/http/api"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/sentry"
	"github.com/GoldenOwlAsia/go-golang-api/internal/i18n"
)

func main() {
	if err := i18n.LoadTranslations(); err != nil {
		panic("Failed to load translations: " + err.Error())
}

	sentry.Init()
	server, err := api.NewServer()

	if err != nil {
		panic("Error when start server")
	}
	server.Run()
}
