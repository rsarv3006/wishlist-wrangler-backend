package dto

import (
	"wishlist-wrangler-api/ent/wishlistsection"

	"github.com/google/uuid"
)

type CreateWishlistDto struct {
	Title      string                      `json:"title"`
	TemplateId uuid.UUID                   `json:"templateId"`
	CreatorId  string                      `json:"creatorId"`
	Sections   []*CreateWishlistSectionDto `json:"sections"`
}

type CreateWishlistSectionDto struct {
	Type              wishlistsection.Type `json:"type"`
	Value             string               `json:"Value"`
	TemplateSectionId uuid.UUID            `json:"templateSectionId"`
}
