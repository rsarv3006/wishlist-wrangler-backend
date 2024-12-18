package usershareddao

import coreshareddao "wishlist-wrangler-api/src/core-shared/core-shared-dao"

type UserEntityRole string

type UserEntity struct {
	coreshareddao.IdentifiedEntity
	Username string `gorm:"column:username;unique;not null" json:"username"`
	Email    string `gorm:"column:email;unique;not null" json:"email"`
	Status   string `gorm:"column:status;not null" json:"status"`
}

func (u *UserEntity) TableName() string {
	return "user"
}

type UserEntityDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

const (
	UserStatusActive                   = "ACTIVE"
	UserStatusEmailVerificationPending = "PENDING"
	UserStatusInactive                 = "DELETED"
)
