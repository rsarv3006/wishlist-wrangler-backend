package dao

import "github.com/google/uuid"

type LoginRequestStatus string

const (
	LoginRequestStatusPending   = "PENDING"
	LoginRequestStatusCompleted = "COMPLETED"
	LoginRequestStatusExpired   = "EXPIRED"
)

type LoginRequest struct {
	IdentifiedEntity
	UserID           uuid.UUID          `gorm:"column:user_id"`
	Email            string             `gorm:"column:email"`
	LoginRequestCode string             `gorm:"column:login_request_code"`
	Status           LoginRequestStatus `gorm:"column:status"`
}
