package main

import (
	"github.com/ldhoax/SOS-requester/http/api"
	"github.com/ldhoax/SOS-requester/pkg/sentry"
	"github.com/ldhoax/SOS-requester/internal/i18n"
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
