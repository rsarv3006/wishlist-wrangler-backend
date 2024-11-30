package service

import (
	"wishlist-wrangler-api/dao"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"

	"github.com/google/uuid"
)

type WishlistTemplateService struct {
	*CrudService[dao.WishlistTemplate, uuid.UUID, dto.CreateWishlistTemplateDto]
}

func NewWishlistTemplateService(repository *repository.CrudRepository[dao.WishlistTemplate, uuid.UUID]) *WishlistTemplateService {
	return &WishlistTemplateService{
		CrudService: NewCrudService[dao.WishlistTemplate, uuid.UUID, dto.CreateWishlistTemplateDto](repository, logger.NewStandardLoggerFactory().CreateLogger("WishlistTemplateService")),
	}
}
