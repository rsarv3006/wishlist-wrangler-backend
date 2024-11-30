package dao

const (
	UserStatusActive                   = "ACTIVE"
	UserStatusEmailVerificationPending = "PENDING"
	UserStatusInactive                 = "DELETED"
)

type User struct {
	IdentifiedEntity
	DisplayName string `gorm:"column:display_name"`
	Email       string `gorm:"column:email"`
	Status      string `gorm:"column:status"`
}
