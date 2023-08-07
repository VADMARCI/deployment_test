package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Cars struct {
	ent.Schema
}

func (Cars) Fields() []ent.Field {
	return []ent.Field{field.Int("id").Unique().Immutable().Annotations(entgql.Skip(entgql.SkipWhereInput))}
}
func (Cars) Edges() []ent.Edge {
	return []ent.Edge{edge.To("dealership_cars", Dealership.Type).Annotations(entgql.Skip(entgql.SkipAll))}
}
func (Cars) Annotations() []schema.Annotation {
	return []schema.Annotation{entgql.QueryField()}
}
func (Cars) Indexes() []ent.Index {
	return []ent.Index{index.Fields("id")}
}
