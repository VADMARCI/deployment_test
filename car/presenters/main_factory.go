package presenters

import (
	"car/gateways"
	"car/gateways/error_gateway"
	"car/gateways/sql_db"
	"context"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/labstack/echo/v4"

	"github.com/pepusz/go_redirect/gateways/auth"
	"github.com/pepusz/go_redirect/gateways/hash"
	"github.com/pepusz/go_redirect/gateways/nats"
)

// TODO: remove this logic. Somehow Middleware function called 5 times during initialization and ent hooks would be added 5 times
var entHooksAdded bool

type MainFactory struct {
	DB          *sql_db.SqlDb
	Gateways    gateways.Gateways
	KetoClient  *auth.KetoClient
	NatsGateway *nats.Gateway
	Redis       *redis.Client
}

func (mf *MainFactory) InitFactory() MainFactory {
	newMainFactory := MainFactory{
		DB:          mf.DB,
		NatsGateway: mf.NatsGateway,
		KetoClient:  mf.KetoClient,
		Redis:       mf.Redis,
	}
	newMainFactory.Gateways = gateways.Gateways{
		ErrorGateway: &error_gateway.Errors{},
		HashGateway:  &hash.Hash{},
	}
	newMainFactory.SetEntHooks()
	return newMainFactory
}

func (mf *MainFactory) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newMainFactory := mf.InitFactory()
		cc := &Context{c}
		cc.SetMainFactory(&newMainFactory)
		ctx := context.WithValue(c.Request().Context(), "echoContext", cc)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(cc)
	}
}

func (mf *MainFactory) SetEntHooks() {
	if !entHooksAdded {
		entHooksAdded = true
		log.Println("Ent hooks added to ent client")
	}
}
