package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	WishlistSectionTypeText  = "TEXT"
	WishlistSectionTypeImage = "IMAGE"
	WishlistSectionTypeVideo = "VIDEO"
	WishlistSectionTypeLink  = "LINK"
)

var WishlistSectionTypeVariant = []string{
	WishlistSectionTypeText,
	WishlistSectionTypeImage,
	WishlistSectionTypeVideo,
	WishlistSectionTypeLink,
}

// WishlistSection holds the schema definition for the WishlistSection entity.
type WishlistSection struct {
	ent.Schema
}

// Fields of the WishlistSection.
func (WishlistSection) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Enum("type").
			Values(WishlistSectionTypeVariant...),
		field.String("value").
			MaxLen(6144).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.UUID("wishlist_id", uuid.UUID{}),
		field.UUID("template_section_id", uuid.UUID{}),
	}
}

// Edges of the WishlistSection.
func (WishlistSection) Edges() []ent.Edge {
	return nil
}
