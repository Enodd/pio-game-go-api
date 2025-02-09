package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("username").
			MinLen(8).
			MaxLen(40).
			NotEmpty().
			Unique(),
		field.String("password").
			MinLen(8).
			MaxLen(40).
			NotEmpty(),
	}
}

// relacje
// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
