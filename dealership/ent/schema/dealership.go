package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Dealership struct {
	ent.Schema
}

func (Dealership) Fields() []ent.Field {
	return []ent.Field{field.String("city"), field.String("name")}
}
func (Dealership) Edges() []ent.Edge {
	return []ent.Edge{edge.From("cars", Cars.Type).Ref("dealership_cars").Annotations(entgql.Skip(entgql.SkipWhereInput))}
}
func (Dealership) Annotations() []schema.Annotation {
	return []schema.Annotation{entgql.QueryField(), entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate())}
}
