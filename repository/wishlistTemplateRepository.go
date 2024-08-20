package repository

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/wishlisttemplate"

	"github.com/google/uuid"
)

func CreateWishlistTemplate(dbClient *ent.Client, templateDto *dto.CreateWishlistTemplateDto, creator *ent.User) (*ent.WishlistTemplate, error) {
	template, err := dbClient.WishlistTemplate.Create().
		SetTitle(templateDto.Title).
		SetDescription(templateDto.Description).
		SetCreatorID(creator.ID).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return dbClient.WishlistTemplate.Query().
		Where(wishlisttemplate.ID(template.ID)).
		Only(context.Background())
}

func GetWishlistTemplate(dbClient *ent.Client, id uuid.UUID) (*ent.WishlistTemplate, error) {
	return dbClient.WishlistTemplate.Query().
		Where(wishlisttemplate.ID(id)).
		Only(context.Background())
}

func GetWishlistTemplatesByUser(dbClient *ent.Client, userId uuid.UUID) ([]*ent.WishlistTemplate, error) {
	return dbClient.WishlistTemplate.
		Query().
		Where(
			wishlisttemplate.And(
				wishlisttemplate.CreatorID(userId),
				wishlisttemplate.StatusEQ(wishlisttemplate.StatusACTIVE),
			),
		).
		All(context.Background())
}

func DeleteWishlistTemplate(dbClient *ent.Client, templateId uuid.UUID, currentUserId uuid.UUID) error {
	template, err := dbClient.WishlistTemplate.Query().
		Where(
			wishlisttemplate.And(
				wishlisttemplate.ID(templateId),
				wishlisttemplate.CreatorID(currentUserId),
			)).
		Only(context.Background())

	if err != nil {
		return err
	}

	if err := dbClient.WishlistTemplate.DeleteOneID(template.ID).Exec(context.Background()); err != nil {
		return err
	}

	return nil
}
