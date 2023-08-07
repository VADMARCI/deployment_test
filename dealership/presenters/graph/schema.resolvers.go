package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dealership/ent"
	"dealership/models"
	"dealership/presenters/graph/generated"
	"dealership/presenters/graph/genresolvers"
)

//THIS IS A TEMPLATE FOR EDIT

func (r *dealershipResolver) Cars(ctx context.Context, obj *ent.Dealership) ([]*models.Car, error) {
	return genresolvers.DealershipCars(ctx, obj)
}

func (r *mutationResolver) CreateDealership(ctx context.Context, input ent.CreateDealershipInput) (*ent.Dealership, error) {
	return genresolvers.CreateDealership(ctx, input)
}

func (r *mutationResolver) DeleteDealership(ctx context.Context, id int) (*bool, error) {
	return genresolvers.DeleteDealership(ctx, id)
}

func (r *queryResolver) Dealership(ctx context.Context, id int) (*ent.Dealership, error) {
	return genresolvers.ReadDealership(ctx, id)
}

func (r *queryResolver) Dealerships(ctx context.Context, input ent.DealershipWhereInput) ([]*ent.Dealership, error) {
	return genresolvers.ListDealership(ctx, input)
}

// Dealership returns generated.DealershipResolver implementation.
func (r *Resolver) Dealership() generated.DealershipResolver { return &dealershipResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type dealershipResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
