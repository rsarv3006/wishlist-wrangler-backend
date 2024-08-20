package dto

import "github.com/google/uuid"

type CreateWishlistDto struct {
	Title      string    `json:"title"`
	TemplateId uuid.UUID `json:"templateId"`
	CreatorId  string    `json:"creatorId"`
}
