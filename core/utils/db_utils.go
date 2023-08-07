package utils

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

const SkipReplicate = "skipReplicate"

func SkipReplicateIfParentIsMutation(ctx context.Context) context.Context {
	rctx := graphql.GetFieldContext(ctx)
	if rctx.Parent != nil && rctx.Parent.Object == "Mutation" {
		return SkipReplication(ctx)
	}
	return ctx
}

func SkipReplication(ctx context.Context) context.Context {
	return context.WithValue(ctx, SkipReplicate, true)
}
