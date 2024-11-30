package service

import (
	"wishlist-wrangler-api/dao"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"

	"github.com/google/uuid"
)

type UserService struct {
	*CrudService[dao.User, uuid.UUID, dto.CreateUserDto]
}

func NewUserService(repository *repository.CrudRepository[dao.User, uuid.UUID]) *UserService {
	return &UserService{
		CrudService: NewCrudService[dao.User, uuid.UUID, dto.CreateUserDto](repository, logger.NewStandardLoggerFactory().CreateLogger("UserService")),
	}
}
