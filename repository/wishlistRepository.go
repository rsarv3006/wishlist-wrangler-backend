package repository

import (
	"context"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/ent/wishlist"

	"github.com/google/uuid"
)

func GetWishlistsByTemplate(dbClient *ent.Client, templateId uuid.UUID) ([]*ent.Wishlist, error) {
	return dbClient.Wishlist.
		Query().
		Where(
			wishlist.TemplateID(templateId),
		).
		All(context.Background())
}
