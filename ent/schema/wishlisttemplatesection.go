package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	WishlistTemplateSectionTypeText  = "TEXT"
	WishlistTemplateSectionTypeImage = "IMAGE"
	WishlistTemplateSectionTypeVideo = "VIDEO"
	WishlistTemplateSectionTypeLink  = "LINK"
)

var WishlistTemplateSectionTypeVariant = []string{
	WishlistTemplateSectionTypeText,
	WishlistTemplateSectionTypeImage,
	WishlistTemplateSectionTypeVideo,
	WishlistTemplateSectionTypeLink,
}

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
		field.UUID("wishlist_template_id", uuid.UUID{}),
		field.String("sectionId").
			MaxLen(255).
			NotEmpty(),
		field.Enum("type").
			Values(WishlistTemplateSectionTypeVariant...),
		field.Int("sort_order").
			NonNegative(),
	}
}

// Edges of the WishlistTemplateSection.
func (WishlistTemplateSection) Edges() []ent.Edge {
	return nil
}
