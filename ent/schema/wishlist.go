package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	WishlistStatusPending   = "PENDING"
	WishlistStatusActive    = "ACTIVE"
	WishlistStatusRemoved   = "REMOVED"
	WishlistStatusCompleted = "COMPLETED"
)

var WishlistStatus = []string{
	WishlistStatusPending,
	WishlistStatusActive,
	WishlistStatusRemoved,
	WishlistStatusCompleted,
}

// Wishlist holds the schema definition for the Wishlist entity.
type Wishlist struct {
	ent.Schema
}

// Fields of the Wishlist.
func (Wishlist) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").
			MaxLen(255).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			Values(WishlistStatus...).
			Default(WishlistStatusPending),
		field.UUID("template_id", uuid.UUID{}),
		field.UUID("creator_id", uuid.UUID{}),
	}
}

// Edges of the Wishlist.
func (Wishlist) Edges() []ent.Edge {
	return nil
}
