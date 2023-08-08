package graph_admin

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/pepusz/go_redirect/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"car/presenters/graph_admin/generated"

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
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {

		res, err = next(utils.SkipReplicateIfParentIsMutation(ctx))

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
	h := playground.Handler("GraphQL", "/query")
	h.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

func Router(e *echo.Echo) {
	e.GET("/", playgroundHandler)
	e.POST("/query", graphqlHandler)
}
