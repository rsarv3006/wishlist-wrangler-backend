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

func GetWishlistById(dbClient *ent.Client, id uuid.UUID) (*ent.Wishlist, error) {
	return dbClient.Wishlist.Query().
		Where(wishlist.ID(id)).
		Only(context.Background())
}

func DeleteWishlist(dbClient *ent.Client, wishlistId uuid.UUID, userId uuid.UUID) error {
	wishlistToDelete, err := dbClient.Wishlist.Query().
		Where(
			wishlist.And(
				wishlist.ID(wishlistId),
				wishlist.CreatorID(userId),
			)).
		Only(context.Background())

	if err != nil {
		return err
	}

	if err := dbClient.Wishlist.DeleteOneID(wishlistToDelete.ID).Exec(context.Background()); err != nil {
		return err
	}

	return nil
}
