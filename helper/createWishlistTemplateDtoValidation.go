package helper

import (
	"wishlist-wrangler-api/dto"
)

func DidCreateWishlistTemplateDtoPassValidation(dto dto.CreateWishlistTemplateDto) bool {
	valid := true
	if len(dto.Title) == 0 {
		valid = false
	}

	if len(dto.TemplateSections) == 0 {
		valid = false
	}

	for _, section := range dto.TemplateSections {
		if len(section.Title) == 0 {
			valid = false
		}

		switch section.Type {
		case "TEXT":
			if len(section.SectionId) == 0 {
				valid = false
			}
		}

		if len(section.Type) == 0 {
			valid = false
		}

	}

	return valid
}
