package service

import (
	"wishlist-wrangler-api/dao"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"

	"github.com/google/uuid"
)

type WishlistTemplateSectionService struct {
	*CrudService[dao.WishlistTemplateSection, uuid.UUID, dto.CreateWishtlistTemplateSectionDto]
}

func NewWishlistTemplateSectionService(repository *repository.CrudRepository[dao.WishlistTemplateSection, uuid.UUID]) *WishlistTemplateSectionService {
	return &WishlistTemplateSectionService{
		CrudService: NewCrudService[dao.WishlistTemplateSection, uuid.UUID, dto.CreateWishtlistTemplateSectionDto](repository, logger.NewStandardLoggerFactory().CreateLogger("WishlistTemplateSectionService")),
	}
}
