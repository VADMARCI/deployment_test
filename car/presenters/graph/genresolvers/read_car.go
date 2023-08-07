package genresolvers

import (
	"car/ent"
	"car/presenters"

	"context"
)

// Generated from templates/actions/read.go.tmpl
func ReadCar(ctx context.Context, id int) (*ent.Car, error) {
	context := ctx.Value("echoContext").(*presenters.Context)

	mainFactory := context.GetMainFactory()
	obj, err := mainFactory.DB.EntClient.Car.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
