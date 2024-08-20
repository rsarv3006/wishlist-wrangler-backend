package helper

import (
	"errors"
	"wishlist-wrangler-api/ent/wishlistsection"
)

func ConvertStringToWishlistSectionType(value string) (wishlistsection.Type, error) {
	switch value {
	case "TEXT":
		return wishlistsection.TypeTEXT, nil
	case "IMAGE":
		return wishlistsection.TypeIMAGE, nil
	case "VIDEO":
		return wishlistsection.TypeVIDEO, nil
	case "LINK":
		return wishlistsection.TypeLINK, nil
	default:
		return wishlistsection.TypeTEXT, errors.New("Invalid wishlist section type")
	}
}
