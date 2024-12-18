package loginrequestsharedservice

import (
	coresharedlogger "wishlist-wrangler-api/src/core-shared/core-shared-logger"
	coresharedservice "wishlist-wrangler-api/src/core-shared/core-shared-service"
	loginrequestshareddao "wishlist-wrangler-api/src/login-request-shared/login-request-shared-dao"
	loginrequestsharedrepository "wishlist-wrangler-api/src/login-request-shared/login-request-shared-repository"
	usersharedservice "wishlist-wrangler-api/src/user-shared/user-shared-service"
)

type LoginRequestService struct {
	*coresharedservice.CrudService[loginrequestshareddao.LoginRequestEntity, uint, loginrequestshareddao.LoginRequestEntityDTO]
	UserService *usersharedservice.UserService
}

func NewLoginRequestService(repository *loginrequestsharedrepository.LoginRequestRepository, logger coresharedlogger.Logger, userService *usersharedservice.UserService) *LoginRequestService {
	return &LoginRequestService{
		CrudService: coresharedservice.NewCrudService[loginrequestshareddao.LoginRequestEntity, uint, loginrequestshareddao.LoginRequestEntityDTO](repository.CrudRepository, logger),
		UserService: userService,
	}
}

func (l *LoginRequestService) CreateLoginRequest(loginRequest *loginrequestshareddao.LoginRequestEntityDTO) (*loginrequestshareddao.LoginRequestEntity, error) {

	return l.CrudService.CreateEntityFromDTO(loginRequest)
}
