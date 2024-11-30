package service

import (
	"wishlist-wrangler-api/dao"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"

	"github.com/google/uuid"
)

type WishlistSectionService struct {
	*CrudService[dao.WishlistSection, uuid.UUID, dto.CreateWishlistSectionDto]
}

func NewWishlistSectionService(repository *repository.CrudRepository[dao.WishlistSection, uuid.UUID]) *WishlistSectionService {
	return &WishlistSectionService{
		CrudService: NewCrudService[dao.WishlistSection, uuid.UUID, dto.CreateWishlistSectionDto](repository, logger.NewStandardLoggerFactory().CreateLogger("WishlistSectionService")),
	}
}
