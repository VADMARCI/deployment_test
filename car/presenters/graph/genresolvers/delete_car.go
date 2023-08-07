package genresolvers

import (
	"car/presenters"

	"context"
)

// Generated from templates/actions/delete.go.tmpl
func DeleteCar(ctx context.Context, id int) (*bool, error) {
	var err error
	context := ctx.Value("echoContext").(*presenters.Context)
	mainFactory := context.GetMainFactory()

	_, err = mainFactory.DB.EntClient.Car.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	err = mainFactory.DB.EntClient.Car.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	b := err == nil

	return &b, err
}
