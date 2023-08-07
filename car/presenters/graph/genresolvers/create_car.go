package genresolvers

import (
	"car/ent"
	"car/presenters"
	"context"
)

// Generated from templates/actions/create.go.tmpl
func CreateCar(ctx context.Context, input ent.CreateCarInput) (*ent.Car, error) {
	context := ctx.Value("echoContext").(*presenters.Context)
	mainFactory := context.GetMainFactory()

	client := mainFactory.DB.EntClient

	obj, err := client.Car.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil

}
