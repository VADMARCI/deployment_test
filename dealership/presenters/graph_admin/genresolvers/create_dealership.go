package genresolvers

import (
	"context"
	"dealership/ent"
	"dealership/presenters"
)

// Generated from templates/actions/create.go.tmpl
func CreateDealership(ctx context.Context, input ent.CreateDealershipInput) (*ent.Dealership, error) {
	context := ctx.Value("echoContext").(*presenters.Context)
	mainFactory := context.GetMainFactory()

	client := mainFactory.DB.EntClient

	bulk := make([]*ent.CarsCreate, len(input.CarIDs))
	for i, d := range input.CarIDs {
		bulk[i] = client.Cars.Create().SetID(d)
	}
	_, err := client.Cars.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}

	obj, err := client.Dealership.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil

}
