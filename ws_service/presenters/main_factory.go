package presenters

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/pepusz/go_redirect/gateways/nats"
)

type MainFactory struct {
	ChannelStore    *ChannelStore
	ConnectionStore *ConnectionStore
	NatsGateway     *nats.Gateway
}

func (mf *MainFactory) InitFactory() MainFactory {
	channelStore := mf.ChannelStore
	if channelStore == nil {
		channelStore = &ChannelStore{}
	}
	if mf.ConnectionStore == nil {
		mf.ConnectionStore = &ConnectionStore{}
	}
	newMainFactory := MainFactory{
		ConnectionStore: mf.ConnectionStore,
		NatsGateway:     mf.NatsGateway,
		// Channels: mf.Channels,
	}
	channelStore.MainFactory = &newMainFactory
	newMainFactory.ChannelStore = channelStore

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
