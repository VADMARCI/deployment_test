package genresolvers

import (
	"car/ent"
	"car/presenters"
	"context"
)

// Generated from templates/actions/list.go.tmpl
func ListCar(ctx context.Context, input ent.CarWhereInput) ([]*ent.Car, error) {
	context := ctx.Value("echoContext").(*presenters.Context)

	mainFactory := context.GetMainFactory()
	baseQuery := mainFactory.DB.EntClient.Car.Query()
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
