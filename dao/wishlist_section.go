package dao

type WishlistSection struct {
	IdentifiedEntity
	Value             string `gorm:"column:value"`
	WishlistID        string `gorm:"column:wishlist_id"`
	TemplateSectionID string `gorm:"column:template_section_id"`
}
