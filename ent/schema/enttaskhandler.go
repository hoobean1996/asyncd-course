package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// EntTaskHandler holds the schema definition for the EntTaskHandler entity.
type EntTaskHandler struct {
	ent.Schema
}

// Fields of the EntTaskHandler.
func (EntTaskHandler) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("handler name"),
		field.String("signature").Comment("handler's signature, E.g: (int) -> int"),
		field.Time("created_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the EntTaskHandler.
func (EntTaskHandler) Edges() []ent.Edge {
	return nil
}

func (EntTaskHandler) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
