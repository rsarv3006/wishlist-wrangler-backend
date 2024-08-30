package dto

import (
	"github.com/google/uuid"
)

type CreateWishlistDto struct {
	Title      string                      `json:"title"`
	TemplateId uuid.UUID                   `json:"templateId"`
	Sections   []*CreateWishlistSectionDto `json:"sections"`
}

type CreateWishlistSectionDto struct {
	Value             string    `json:"value"`
	TemplateSectionId uuid.UUID `json:"templateSectionId"`
}
