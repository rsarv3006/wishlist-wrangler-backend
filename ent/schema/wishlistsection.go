package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	WishlistSectionTypeText    = "WISHLIST_SECTION_TYPE_TEXT"
	WishlistSectionTypeImage   = "WISHLIST_SECTION_TYPE_IMAGE"
	WishlistSectionTypeVideo   = "WISHLIST_SECTION_TYPE_VIDEO"
	WishlistSectionTypeLink    = "WISHLIST_SECTION_TYPE_LINK"
	WishlistSectionTypeBoolean = "WISHLIST_SECTION_TYPE_BOOLEAN"
)

var WishlistSectionTypeVariant = []string{
	WishlistSectionTypeText,
	WishlistSectionTypeImage,
	WishlistSectionTypeVideo,
	WishlistSectionTypeLink,
	WishlistSectionTypeBoolean,
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
		field.String("textValue").
			MaxLen(1024).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the WishlistSection.
func (WishlistSection) Edges() []ent.Edge {
	return nil
}
