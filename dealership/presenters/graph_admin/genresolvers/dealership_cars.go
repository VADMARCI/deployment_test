package genresolvers

import (
	"context"
	"dealership/ent"
	"dealership/models_admin"
)

// Generated from templates/resolvers/external_reference_list_value.go.tmpl
func DealershipCars(ctx context.Context, obj *ent.Dealership) ([]*models_admin.Car, error) {
	objects, err := obj.QueryCars().All(ctx)
	if err != nil {
		return nil, nil
	}
	var values []*models_admin.Car

	for _, object := range objects {
		values = append(values, &models_admin.Car{ID: object.ID})

	}

	return values, nil
}
