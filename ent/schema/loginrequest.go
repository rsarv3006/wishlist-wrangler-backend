package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	LoginRequestStatusPending   = "PENDING"
	LoginRequestStatusCompleted = "COMPLETED"
	LoginRequestStatusExpired   = "EXPIRED"
)

var LoginRequestStatus = []string{
	LoginRequestStatusPending,
	LoginRequestStatusCompleted,
	LoginRequestStatusExpired,
}

// LoginRequest holds the schema definition for the LoginRequest entity.
type LoginRequest struct {
	ent.Schema
}

// Fields of the LoginRequest.
func (LoginRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("userId", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.String("loginRequestCode"),
		field.Enum("status").
			Values(LoginRequestStatus...).
			Default(LoginRequestStatusPending),
	}
}

// Edges of the LoginRequest.
func (LoginRequest) Edges() []ent.Edge {
	return nil
}
