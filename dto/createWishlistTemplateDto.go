package dto

type CreateWishlistTemplateDto struct {
	Title            string                              `json:"name" validate:"required"`
	Description      string                              `json:"description"`
	TemplateSections []CreateWishtlistTemplateSectionDto `json:"templateSections"`
}

type CreateWishtlistTemplateSectionDto struct {
	Title string `json:"title" validate:"required"`
}
