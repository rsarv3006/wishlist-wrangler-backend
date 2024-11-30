package dao

type WishlistTemplateSection struct {
	IdentifiedEntity
	Title              string `gorm:"column:title"`
	SectionId          string `gorm:"column:section_id"`
	Type               string `gorm:"column:type"`
	SortOrder          int    `gorm:"column:sort_order"`
	WishlistTemplateID string `gorm:"column:wishlist_template_id"`
}
