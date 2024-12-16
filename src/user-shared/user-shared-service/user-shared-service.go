package usersharedservice

import (
	coresharedlogger "wishlist-wrangler-api/src/core-shared/core-shared-logger"
	coresharedservice "wishlist-wrangler-api/src/core-shared/core-shared-service"
	usershareddao "wishlist-wrangler-api/src/user-shared/user-shared-dao"
	usersharedrepository "wishlist-wrangler-api/src/user-shared/user-shared-repository"
)

type UserService struct {
	*coresharedservice.CrudService[usershareddao.UserEntity, uint, usershareddao.UserEntityDTO]
}

func NewUserService(repository *usersharedrepository.UserRepository, logger coresharedlogger.Logger) *UserService {
	return &UserService{
		CrudService: coresharedservice.NewCrudService[usershareddao.UserEntity, uint, usershareddao.UserEntityDTO](repository.CrudRepository, logger),
	}
}
