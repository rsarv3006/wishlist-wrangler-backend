package repository

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
)

func CreateWishlistTemplateSections(dbClient *ent.Client, templateSectionDto []dto.CreateWishtlistTemplateSectionDto) ([]*ent.WishlistTemplateSection, error) {
	var sections []*ent.WishlistTemplateSection = make([]*ent.WishlistTemplateSection, len(templateSectionDto))

	for _, sectionDto := range templateSectionDto {
		transaction, err := dbClient.Tx(context.Background())

		if err != nil {
			return nil, err
		}
		defer transaction.Rollback()

		section, err := transaction.WishlistTemplateSection.Create().
			SetTitle(sectionDto.Title).
			Save(context.Background())

		if err != nil {
			return nil, err
		}

		sections = append(sections, section)
	}

	return nil, nil
}
