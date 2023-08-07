package genresolvers

import (
	"dealership/ent"
	"dealership/presenters"

	"context"
)

// Generated from templates/actions/read.go.tmpl
func ReadDealership(ctx context.Context, id int) (*ent.Dealership, error) {
	context := ctx.Value("echoContext").(*presenters.Context)

	mainFactory := context.GetMainFactory()
	obj, err := mainFactory.DB.EntClient.Dealership.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
