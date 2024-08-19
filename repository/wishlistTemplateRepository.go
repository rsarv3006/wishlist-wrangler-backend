package repository

import (
	"context"
	"log"
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
		log.Println("failed to create wishlist template - you know where you are")

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
				// wishlisttemplate.HasCreatorWith(user.ID(userId)),
				wishlisttemplate.StatusEQ(wishlisttemplate.StatusACTIVE),
			),
		).
		All(context.Background())
}
