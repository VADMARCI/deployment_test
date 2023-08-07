package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
	"ws_service/presenters"
	"ws_service/presenters/consumers"
	"ws_service/presenters/graph"
	middleware2 "ws_service/presenters/middleware"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pepusz/go_redirect/gateways/nats"
)

// var channels = presenters.NewChannels()

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DNS"),
		AttachStacktrace: true,
		Environment:      os.Getenv("ENV"),
		Release:          fmt.Sprint(os.Getenv("BRANCH"), "(", os.Getenv("COMMIT"), ")"),
	})
	if err != nil {
		log.Printf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	defer sentry.Recover()

	natsGateway := nats.NewGateway()

	mainFactory := presenters.MainFactory{
		// Channels: &channels,
		NatsGateway: &natsGateway,
	}
	mainFactory = mainFactory.InitFactory()

	defer mainFactory.NatsGateway.Stop()

	e := echo.New()
	e.Use(mainFactory.Middleware)
	e.Use(middleware2.JWT)
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		re := regexp.MustCompile(`"password":.+(\s|\t|\n|\})`)
		reqBodyString := re.ReplaceAllString(string(reqBody), "password: \\\"[FILTERED]\\\"")
		log.Println(reqBodyString)
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: middleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: []string{"*"}, // []string{http.MethodGet, http.MethodHead,
		// http.MethodPut, http.MethodPatch, http.MethodPost,
		// http.MethodDelete, http.MethodOptions},
		ExposeHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlExposeHeaders,
		},
	}))
	graph.Router(e)

	consumers.InitConsumers(&mainFactory)
	e.Logger.Fatal(e.Start(":" + port))
}
