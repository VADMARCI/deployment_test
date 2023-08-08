package graph_admin

import (
	"context"
	"dealership/presenters"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func ResponseHandler(ctx context.Context) error {
	context := ctx.Value("echoContext").(*presenters.Context)
	mainFactory := context.GetMainFactory()
	if mainFactory.Gateways.ErrorGateway.HasError() {
		for _, error := range mainFactory.Gateways.ErrorGateway.GetErrors() {
			graphql.AddError(ctx, &gqlerror.Error{
				Message:    error.Message,
				Extensions: error.Extensions,
			})
		}
		return fmt.Errorf("has errors")
	}
	return nil
}
