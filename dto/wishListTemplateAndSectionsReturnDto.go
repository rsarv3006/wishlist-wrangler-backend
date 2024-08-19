package dto

import (
	"wishlist-wrangler-api/ent"
)

type WishlistTemplateAndSectionsReturnDto struct {
	WishlistTemplate         *ent.WishlistTemplate          `json:"wishlistTemplate"`
	WishlistTemplateSections []*ent.WishlistTemplateSection `json:"wishlistTemplateSections"`
}
