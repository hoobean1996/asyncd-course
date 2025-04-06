package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// EntTask holds the schema definition for the EntTask entity.
type EntTask struct {
	ent.Schema
}

// Fields of the EntTask.
func (EntTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("handler").Comment("This is the handler of the task"),
		field.String("parameter").Comment("This is JSON input of the task"),
		field.Int("priority").Comment("This is the priority of the task"),
		field.Time("created_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the EntTask.
func (EntTask) Edges() []ent.Edge {
	return nil
}

func (EntTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
