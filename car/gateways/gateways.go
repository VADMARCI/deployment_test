package gateways

import (
	"github.com/pepusz/go_redirect/gateways"
	"github.com/pepusz/go_redirect/models"
)

type Gateways struct {
	ErrorGateway ErrorsGatewayInterface
	HashGateway  gateways.HashGatewayInterface
}

type ErrorsGatewayInterface interface {
	AddError(error models.Error)
	GetErrors() []models.Error
	HasError() bool
	Clear()
}
