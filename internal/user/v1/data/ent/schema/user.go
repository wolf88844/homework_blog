package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.Uint32("age"),
		field.Uint32("sex"),
		field.Time("create_at").Default(time.Now).SchemaType(map[string]string{dialect.MySQL:"datetime"}),
		field.Time("update_at").Default(time.Now).SchemaType(map[string]string{dialect.MySQL:"datetime"}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
