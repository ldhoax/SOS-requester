package sentry

import (
	"fmt"
	"log"
	"time"

	"github.com/GoldenOwlAsia/go-golang-api/configs"
	"github.com/getsentry/sentry-go"
)

func Init() {
	cnf, err := configs.NewParsedConfig()
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              cnf.Sentry.Dsn,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func Log(issue error) {
	Init()
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureException(issue)
}
