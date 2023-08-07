package graph

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"ws_service/presenters"
	"ws_service/presenters/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	core_models "github.com/pepusz/go_redirect/models"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"os"
	"runtime/debug"
)

func graphqlHandler(c echo.Context) error {

	resolverConfig := generated.Config{Resolvers: &Resolver{}}
	resolverConfig.Directives.HooResolver = func(
		ctx context.Context, obj interface{}, next graphql.Resolver, action string,
	) (res interface{}, err error) {
		return next(ctx)
	}

	srv := handler.New(generated.NewExecutableSchema(resolverConfig))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			cc := c.(*presenters.Context)

			if len(initPayload) == 0 {
				return ctx, nil
			}

			clientId := initPayload.GetString("clientID")
			cc.SetClientId(clientId)
			logrus.Println(clientId)

			r := cc.Request()

			authServiceURL := os.Getenv("AUTH_URL")

			// Create new HTTP request
			newRequest, err := http.NewRequest("GET", fmt.Sprintf("%s/get_user_jwt_token", authServiceURL), nil)
			if err != nil {
				log.Fatalln(err)
			}
			origin := initPayload.GetString("Origin")
			if origin == "" {
				origin = r.Header.Get("Origin")
			}

			if r.Header.Get("Authorization") == "" && len(initPayload.GetString("Authorization")) > 0 {
				r.Header.Set("Authorization", initPayload.GetString("Authorization"))
			}
			// Copy the headers from the original request.
			for name, values := range r.Header {
				for _, value := range values {
					newRequest.Header.Add(name, value)
				}
			}

			crypto := sha256.Sum256([]byte(origin))
			// Copy the cookies from the original request.
			for _, cookie := range r.Cookies() {
				if strings.Contains(cookie.Name, fmt.Sprintf("_%x", crypto)) {
					logrus.Infof("Cookie found %s", cookie.Name)
					cookie.Name = strings.TrimSuffix(cookie.Name, fmt.Sprintf("_%x", crypto))
					newRequest.AddCookie(cookie)
				}
			}
			client := &http.Client{}
			resp, err := client.Do(newRequest)
			if err != nil {
				log.Fatalln(err)
			}

			if resp.StatusCode == http.StatusOK {
				var authResp struct {
					ServiceJWTToken string `json:"service_jwt_token"`
				}

				bodyBytes, err := io.ReadAll(resp.Body)
				err = json.Unmarshal(bodyBytes, &authResp)

				if err != nil {
					log.Println(err)
					return ctx, nil
				}
				fmt.Printf("JWT token got from server.js: %s\n", authResp.ServiceJWTToken)
				user, err := core_models.GetCoreUserFromToken(authResp.ServiceJWTToken)
				if err != nil {
					fmt.Printf("WARN: service-jwt-token header is missing or not valid\n")
				}
				cc.SetCurrentUser(user)
			}

			return ctx, nil
		},
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Subprotocols:    []string{"binary", "graphql-ws"},
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		res, err = next(ctx)
		return res, err
	})

	srv.SetQueryCache(lru.New(1000))
	srv.SetRecoverFunc(graphql.RecoverFunc(func(ctx context.Context, err interface{}) error {
		sentry.CaptureException(fmt.Errorf(fmt.Sprintf("Internal server error.go! %v", err)))

		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr)
		debug.PrintStack()
		return fmt.Errorf("Internal server error!")
	}))

	srv.SetErrorPresenter(
		func(ctx context.Context, e error) *gqlerror.Error {
			// TODO any special logic you want to do here. Must specify path for correct null bubbling behaviour.

			return graphql.DefaultErrorPresenter(ctx, e)
		},
	)
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.ServeHTTP(c.Response(), c.Request())
	return nil
}

func playgroundHandler(c echo.Context) error {
	h := playground.Handler("GraphQL", "/ws-query")
	h.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

func Router(e *echo.Echo) {
	e.GET("/ws-query/playground", playgroundHandler)
	e.GET("/.well-known/apollo/server-health", playgroundHandler)
	e.POST("/ws-query", graphqlHandler)
	e.GET("/ws-query", graphqlHandler)

}
