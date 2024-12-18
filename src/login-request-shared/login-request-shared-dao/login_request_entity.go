package loginrequestshareddao

import coreshareddao "wishlist-wrangler-api/src/core-shared/core-shared-dao"

type LoginRequestEntity struct {
	coreshareddao.IdentifiedEntity
	UserId           int    `gorm:"column:user_id;not null" json:"userId"`
	Email            string `gorm:"column:email;not null" json:"email"`
	LoginRequestCode string `gorm:"column:login_request_code;not null" json:"loginRequestCode"`
	Status           string `gorm:"column:status;not null" json:"status"`
}

func (l *LoginRequestEntity) TableName() string {
	return "login_request"
}

type LoginRequestEntityDTO struct {
	UserId           int    `json:"userId"`
	Email            string `json:"email"`
	LoginRequestCode string `json:"loginRequestCode"`
	Status           string `json:"status"`
}

const (
	LoginRequestStatusPending   = "PENDING"
	LoginRequestStatusCompleted = "COMPLETED"
	LoginRequestStatusExpired   = "EXPIRED"
)
