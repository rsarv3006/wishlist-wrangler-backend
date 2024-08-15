package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	UserStatusActive                   = "ACTIVE"
	UserStatusEmailVerificationPending = "PENDING"
	UserStatusInactive                 = "DELETED"
)

var UserStatus = []string{
	UserStatusActive,
	UserStatusEmailVerificationPending,
	UserStatusInactive,
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("displayName").
			MaxLen(255).
			NotEmpty(),
		field.String("email").
			MaxLen(255).
			NotEmpty().
			Unique(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			Values(UserStatus...).
			Default(UserStatusEmailVerificationPending),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
