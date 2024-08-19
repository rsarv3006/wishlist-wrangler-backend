package repository

import (
	"context"
	"log"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"

	"github.com/google/uuid"
)

func CreateWishlistTemplateSections(dbClient *ent.Client,
	templateSectionDto []dto.CreateWishtlistTemplateSectionDto,
	templateId uuid.UUID) ([]*ent.WishlistTemplateSection, error) {

	sections := make([]*ent.WishlistTemplateSection, 0, len(templateSectionDto))
	// TODO: Double creating the sections
	for _, sectionDto := range templateSectionDto {

		section, err := dbClient.WishlistTemplateSection.Create().
			SetTitle(sectionDto.Title).
			SetWishlistTemplateID(templateId).
			SetType(sectionDto.Type).
			SetSectionId(sectionDto.SectionId).
			Save(context.Background())

		if err != nil {
			return nil, err
		}

		if section == nil {
			return nil, err
		}

		sections = append(sections, section)
	}
	log.Println("sections")
	log.Println(sections)

	return sections, nil
}
