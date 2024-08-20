package repository

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/wishlisttemplatesection"

	"github.com/google/uuid"
)

func CreateWishlistTemplateSections(
	dbClient *ent.Client,
	templateSectionDto []dto.CreateWishtlistTemplateSectionDto,
	templateId uuid.UUID,
) ([]*ent.WishlistTemplateSection, error) {
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

func DeleteWishlistTemplateSection(ctx context.Context, dbClient *ent.Client, sectionId uuid.UUID) error {
	return dbClient.WishlistTemplateSection.DeleteOneID(sectionId).Exec(ctx)
}

func DeleteWishlistTemplateSections(ctx context.Context, dbClient *ent.Client, sectionIds []uuid.UUID) error {
	if len(sectionIds) == 0 {
		return nil
	}

	_, err := dbClient.WishlistTemplateSection.Delete().
		Where(wishlisttemplatesection.IDIn(sectionIds...)).
		Exec(ctx)

	return err
}

func DeleteTemplateSectionsForTemplate(ctx context.Context, dbClient *ent.Client, templateId uuid.UUID) error {
	_, err := dbClient.WishlistTemplateSection.Delete().
		Where(wishlisttemplatesection.WishlistTemplateID(templateId)).
		Exec(ctx)

	return err
}
