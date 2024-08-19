package dto

type CreateWishlistTemplateDto struct {
	Title            string                              `json:"title" validate:"required"`
	Description      string                              `json:"description"`
	TemplateSections []CreateWishtlistTemplateSectionDto `json:"templateSections"`
}

type CreateWishtlistTemplateSectionDto struct {
	Title     string `json:"title" validate:"required"`
	Type      string `json:"type" validate:"required"`
	SectionId string `json:"sectionId" validate:"required"`
}
