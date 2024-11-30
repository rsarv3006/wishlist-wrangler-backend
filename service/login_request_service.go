package service

import (
	"wishlist-wrangler-api/dao"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"

	"github.com/google/uuid"
)

type LoginRequestService struct {
	*CrudService[dao.LoginRequest, uuid.UUID, dto.CreateLoginRequest]
}

func NewLoginRequestService(repository *repository.CrudRepository[dao.LoginRequest, uuid.UUID]) *LoginRequestService {
	return &LoginRequestService{
		CrudService: NewCrudService[dao.LoginRequest, uuid.UUID, dto.CreateLoginRequest](repository, logger.NewStandardLoggerFactory().CreateLogger("LoginRequestService")),
	}
}
