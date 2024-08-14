package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WishlistTemplateSection holds the schema definition for the WishlistTemplateSection entity.
type WishlistTemplateSection struct {
	ent.Schema
}

// Fields of the WishlistTemplateSection.
func (WishlistTemplateSection) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").
			MaxLen(255).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the WishlistTemplateSection.
func (WishlistTemplateSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("wishlistTemplate", WishlistTemplate.Type).
			Ref("sections").
			Unique(),
	}
}
