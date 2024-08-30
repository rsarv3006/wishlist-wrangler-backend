package dto

import "wishlist-wrangler-api/ent/wishlisttemplatesection"

type CreateWishlistTemplateDto struct {
	Title            string                              `json:"title" validate:"required"`
	Description      string                              `json:"description"`
	TemplateSections []CreateWishtlistTemplateSectionDto `json:"templateSections"`
}

type CreateWishtlistTemplateSectionDto struct {
	Title     string                       `json:"title" validate:"required"`
	Type      wishlisttemplatesection.Type `json:"type"`
	SectionId string                       `json:"sectionId" validate:"required"`
}
