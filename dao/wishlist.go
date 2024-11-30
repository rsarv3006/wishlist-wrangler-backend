package dao

type Wishlist struct {
	IdentifiedEntity
	Title      string `gorm:"column:title"`
	TemplateID string `gorm:"column:template_id"`
	CreatorID  string `gorm:"column:creator_id"`
	Status     string `gorm:"column:status"`
}
