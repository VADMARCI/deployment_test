package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pepusz/go_redirect/gateways/nats"

	_ "car/ent/runtime"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"

	"github.com/pepusz/go_redirect/utils"
	//"github.com/pepusz/go_redirect/messaging"
	"car/buckets"
	_ "car/ent/runtime"
	"car/gateways/sql_db"
	"car/presenters"
	"car/presenters/graph_admin"
	middleware2 "car/presenters/middleware"
	"car/resolvers"
	"car/streams"

	"github.com/pepusz/go_redirect/gateways/auth"

	"time"
)

//TODO separate migration from sql_db.go

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	logrus.Info("Starting car service for: admin API")
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
	buckets.InitBuckets()
	db := sql_db.SqlDb{}
	db.Connect()
	devMode := utils.GetEnvBool("DEV")
	if devMode {
		db.Migrate()
	}
	defer db.Close()
	natsGateway := nats.NewGateway()
	ketoClient := auth.NewKetoClient()
	redis := goredislib.NewClient(&goredislib.Options{
		Addr: utils.GetEnvString("REDIS_HOST"),
	})
	defer redis.Close()
	mainFactory := presenters.MainFactory{
		DB:          &db,
		NatsGateway: &natsGateway,
		KetoClient:  ketoClient,
		Redis:       redis,
	}
	defer mainFactory.NatsGateway.Stop()
	e := echo.New()
	e.Use(mainFactory.Middleware)
	e.Use(middleware2.User)
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		re := regexp.MustCompile(`"password":.+(\s|\t|\n|\})`)
		reqBodyString := re.ReplaceAllString(string(reqBody), "password: \\\"[FILTERED]\\\"")
		log.Println(reqBodyString)
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: middleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodHead,
			http.MethodPut, http.MethodPatch, http.MethodPost,
			http.MethodDelete, http.MethodOptions},
		ExposeHeaders: []string{echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlExposeHeaders},
	}))
	graph_admin.Router(e)
	resolvers.CustomRest(e)
	streams.CreateStreams(mainFactory.NatsGateway)
	e.Logger.Fatal(e.Start(":" + port))
}
