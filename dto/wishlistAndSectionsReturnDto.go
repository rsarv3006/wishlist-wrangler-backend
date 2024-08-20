package dto

import (
	"wishlist-wrangler-api/ent"
)

type WishlistAndSectionsReturnDto struct {
	Wishlist         *ent.Wishlist          `json:"wishlist"`
	WishlistSections []*ent.WishlistSection `json:"wishlistSections"`
}
