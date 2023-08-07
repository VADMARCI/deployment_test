package genresolvers

import (
	"car/ent"
	"car/presenters"
	"context"
)

// Generated from templates/resolvers/entity_resolver_by_id.go.tmpl
func FindCarByID(ctx context.Context, id int) (*ent.Car, error) {
	if id == 0 {
		return nil, nil
	}
	context := ctx.Value("echoContext").(*presenters.Context)
	mainFactory := context.GetMainFactory()
	return mainFactory.DB.EntClient.Car.Get(ctx, id)
}
