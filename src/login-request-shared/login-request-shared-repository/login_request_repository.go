package loginrequestsharedrepository

import (
	coresharedrepository "wishlist-wrangler-api/src/core-shared/core-shared-repository"
	loginrequestshareddao "wishlist-wrangler-api/src/login-request-shared/login-request-shared-dao"

	"gorm.io/gorm"
)

type LoginRequestRepository struct {
	*coresharedrepository.CrudRepository[loginrequestshareddao.LoginRequestEntity, uint]
}

func NewLoginRequestRepository(db *gorm.DB) *LoginRequestRepository {
	return &LoginRequestRepository{
		CrudRepository: coresharedrepository.NewCrudRepository[loginrequestshareddao.LoginRequestEntity, uint](db),
	}
}
