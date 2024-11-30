package dao

type WishlistTemplate struct {
	IdentifiedEntity
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Status      string `gorm:"column:status"`
	CreatorID   string `gorm:"column:creator_id"`
}
