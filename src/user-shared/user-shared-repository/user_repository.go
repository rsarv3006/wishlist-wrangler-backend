package usersharedrepository

import (
	coresharedrepository "wishlist-wrangler-api/src/core-shared/core-shared-repository"
	usershareddao "wishlist-wrangler-api/src/user-shared/user-shared-dao"

	"gorm.io/gorm"
)

type UserRepository struct {
	*coresharedrepository.CrudRepository[usershareddao.UserEntity, uint]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		CrudRepository: coresharedrepository.NewCrudRepository[usershareddao.UserEntity, uint](db),
	}
}
