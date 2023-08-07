package genresolvers

import (
	"context"
	"dealership/ent"
	"dealership/presenters"
)

// Generated from templates/actions/list.go.tmpl
func ListDealership(ctx context.Context, input ent.DealershipWhereInput) ([]*ent.Dealership, error) {
	context := ctx.Value("echoContext").(*presenters.Context)

	mainFactory := context.GetMainFactory()
	baseQuery := mainFactory.DB.EntClient.Dealership.Query()
	baseQuery, err := input.Filter(baseQuery)
	if err != nil {
		return nil, err
	}
	obj, err := baseQuery.All(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
