package repository

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/wishlistsection"

	"github.com/google/uuid"
)

func CreateWishlistSection(dbClient *ent.Client, createWishlistSectionDto *dto.CreateWishlistSectionDto, wishlistId uuid.UUID) (*ent.WishlistSection, error) {
	return dbClient.WishlistSection.Create().
		SetType(createWishlistSectionDto.Type).
		SetValue(createWishlistSectionDto.Value).
		SetWishlistID(wishlistId).
		SetTemplateSectionID(createWishlistSectionDto.TemplateSectionId).
		Save(context.Background())
}

func CreateWishlistSections(dbClient *ent.Client, createWishlistSectionDtos []*dto.CreateWishlistSectionDto, wishlistId uuid.UUID) ([]*ent.WishlistSection, error) {
	var sections []*ent.WishlistSection
	for _, createWishlistSectionDto := range createWishlistSectionDtos {
		section, err := CreateWishlistSection(dbClient, createWishlistSectionDto, wishlistId)
		if err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}
	return sections, nil
}

func GetWishlistSectionsForWishlistArray(dbClient *ent.Client, wishlistIds []uuid.UUID) ([]*ent.WishlistSection, error) {
	sections, err := dbClient.WishlistSection.
		Query().
		Where(wishlistsection.WishlistIDIn(wishlistIds...)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func GetWishlistSectionsByWishlistId(dbClient *ent.Client, wishlistId uuid.UUID) ([]*ent.WishlistSection, error) {
	return dbClient.WishlistSection.Query().
		Where(wishlistsection.WishlistID(wishlistId)).
		All(context.Background())
}

func GetWishlistSectionById(dbClient *ent.Client, id uuid.UUID) (*ent.WishlistSection, error) {
	return dbClient.WishlistSection.Query().
		Where(wishlistsection.ID(id)).
		Only(context.Background())
}

func DeleteWishlistSectionsBySectionIds(dbClient *ent.Client, ids []uuid.UUID) error {
	_, err := dbClient.WishlistSection.Delete().
		Where(wishlistsection.IDIn(ids...)).
		Exec(context.Background())
	return err
}

func DeleteWishlistSectionById(dbClient *ent.Client, id uuid.UUID) error {
	_, err := dbClient.WishlistSection.Delete().
		Where(wishlistsection.ID(id)).
		Exec(context.Background())
	return err
}

func DeleteWishlistSectionsByWishlistId(dbClient *ent.Client, wishlistId uuid.UUID) error {
	_, err := dbClient.WishlistSection.Delete().
		Where(wishlistsection.WishlistID(wishlistId)).
		Exec(context.Background())
	return err
}
