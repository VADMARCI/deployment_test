package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Car struct {
	ent.Schema
}

func (Car) Fields() []ent.Field {
	return []ent.Field{field.Bool("is_sold").Default(false), field.String("name"), field.Int("price")}
}
func (Car) Edges() []ent.Edge {
	return nil
}
func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{entgql.QueryField(), entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate())}
}
