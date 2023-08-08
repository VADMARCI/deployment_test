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

func (r *entityResolver) FindCarByID(ctx context.Context, id int) (*ent.Car, error) {
	return genresolvers.FindCarByID(ctx, id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
