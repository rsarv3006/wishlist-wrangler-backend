package usershareddao

import coreshareddao "wishlist-wrangler-api/src/core-shared/core-shared-dao"

type UserEntityRole string

const (
	UserRole  UserEntityRole = "User"
	AdminRole UserEntityRole = "Admin"
)

type UserEntity struct {
	coreshareddao.IdentifiedEntity
	Username    string         `gorm:"column:username;unique;not null" json:"username"`
	Email       string         `gorm:"column:email;unique;not null" json:"email"`
	Role        UserEntityRole `gorm:"column:role;default:User;not null" json:"role"`
	Name        *string        `gorm:"column:name;type:text;" json:"name,omitempty"`
	PhoneNumber *string        `gorm:"column:phone_number;type:text;" json:"phoneNumber,omitempty"`
}

func (u *UserEntity) TableName() string {
	return "user"
}

type UserEntityDTO struct {
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Role        UserEntityRole `json:"role"`
	Name        *string        `json:"name,omitempty"`
	PhoneNumber *string        `json:"phoneNumber,omitempty"`
}
