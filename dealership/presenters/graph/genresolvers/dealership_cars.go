package genresolvers

import (
	"context"
	"dealership/ent"
	"dealership/models"
)

// Generated from templates/resolvers/external_reference_list_value.go.tmpl
func DealershipCars(ctx context.Context, obj *ent.Dealership) ([]*models.Car, error) {
	objects, err := obj.QueryCars().All(ctx)
	if err != nil {
		return nil, nil
	}
	var values []*models.Car

	for _, object := range objects {
		values = append(values, &models.Car{ID: object.ID})

	}

	return values, nil
}
