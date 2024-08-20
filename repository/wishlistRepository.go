package repository

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/wishlist"

	"github.com/google/uuid"
)

func CreateWishlist(dbClient *ent.Client, createWishlistDto dto.CreateWishlistDto, currentUser *ent.User) (*ent.Wishlist, error) {
	return dbClient.Wishlist.
		Create().
		SetTitle(createWishlistDto.Title).
		SetTemplateID(createWishlistDto.TemplateId).
		SetCreatorID(currentUser.ID).
		Save(context.Background())
}

func GetWishlistsByUser(dbClient *ent.Client, userId uuid.UUID) ([]*ent.Wishlist, error) {
	return dbClient.Wishlist.
		Query().
		Where(
			wishlist.CreatorID(userId),
		).
		All(context.Background())
}

func GetWishlistsByTemplate(dbClient *ent.Client, templateId uuid.UUID) ([]*ent.Wishlist, error) {
	return dbClient.Wishlist.
		Query().
		Where(
			wishlist.TemplateID(templateId),
		).
		All(context.Background())
}
