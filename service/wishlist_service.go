package service

import (
	"wishlist-wrangler-api/dao"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"

	"github.com/google/uuid"
)

type WishlistService struct {
	*CrudService[dao.Wishlist, uuid.UUID, dto.CreateWishlistDto]
}

func NewWishlistService(repository *repository.CrudRepository[dao.Wishlist, uuid.UUID]) *WishlistService {
	return &WishlistService{
		CrudService: NewCrudService[dao.Wishlist, uuid.UUID, dto.CreateWishlistDto](repository, logger.NewStandardLoggerFactory().CreateLogger("WishlistService")),
	}
}
