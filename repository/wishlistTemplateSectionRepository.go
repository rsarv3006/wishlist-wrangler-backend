package repository

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/wishlisttemplatesection"

	"github.com/google/uuid"
)

func CreateWishlistTemplateSections(dbClient *ent.Client,
	templateSectionDto []dto.CreateWishtlistTemplateSectionDto,
	templateId uuid.UUID) ([]*ent.WishlistTemplateSection, error) {

	sections := make([]*ent.WishlistTemplateSection, 0, len(templateSectionDto))
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

	return sections, nil
}

func GetWishlistTemplateSections(dbClient *ent.Client, templateId uuid.UUID) ([]*ent.WishlistTemplateSection, error) {
	return dbClient.WishlistTemplateSection.
		Query().
		Where(
			wishlisttemplatesection.And(
				wishlisttemplatesection.WishlistTemplateID(templateId),
			),
		).
		All(context.Background())
}

func GetWishlistTemplateSectionsForTemplateArray(dbClient *ent.Client, templateIds []uuid.UUID) ([]*ent.WishlistTemplateSection, error) {
	sections, err := dbClient.WishlistTemplateSection.
		Query().
		Where(wishlisttemplatesection.WishlistTemplateIDIn(templateIds...)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return sections, nil
}
