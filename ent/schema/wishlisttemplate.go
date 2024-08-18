package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	WishlistTemplateStatusActive  = "ACTIVE"
	WishlistTemplateStatusRemoved = "REMOVED"
)

var WishlistTemplateStatus = []string{
	WishlistTemplateStatusActive,
	WishlistTemplateStatusRemoved,
}

// WishlistTemplate holds the schema definition for the WishlistTemplate entity.
type WishlistTemplate struct {
	ent.Schema
}

// Fields of the WishlistTemplate.
func (WishlistTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").
			MaxLen(255).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.String("description").
			MaxLen(525),
		field.Enum("status").
			Values(WishlistTemplateStatus...).
			Default(WishlistTemplateStatusActive),
	}
}

// Edges of the WishlistTemplate.
func (WishlistTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type),
		edge.To("sections", WishlistTemplateSection.Type),
	}
}
