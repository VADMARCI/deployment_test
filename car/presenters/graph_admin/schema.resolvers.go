package graph_admin

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"car/ent"
	"car/presenters/graph_admin/generated"
	"car/presenters/graph_admin/genresolvers"
	"context"
)

//THIS IS A TEMPLATE FOR EDIT

func (r *mutationResolver) CreateCar(ctx context.Context, input ent.CreateCarInput) (*ent.Car, error) {
	return genresolvers.CreateCar(ctx, input)
}

func (r *mutationResolver) DeleteCar(ctx context.Context, id int) (*bool, error) {
	return genresolvers.DeleteCar(ctx, id)
}

func (r *queryResolver) Car(ctx context.Context, id int) (*ent.Car, error) {
	return genresolvers.ReadCar(ctx, id)
}

func (r *queryResolver) Cars(ctx context.Context, input ent.CarWhereInput) ([]*ent.Car, error) {
	return genresolvers.ListCar(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
