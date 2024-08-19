package helper

import (
	"testing"
	"wishlist-wrangler-api/dto"
)

func TestDidCreateWishlistTemplateDtoPassValidation(t *testing.T) {
	t.Run("Should return true if the dto is valid", func(t *testing.T) {
		dto := dto.CreateWishlistTemplateDto{
			Title:       "Test Title",
			Description: "Test Description",
			TemplateSections: []dto.CreateWishtlistTemplateSectionDto{
				{
					Title:     "Test Section Title",
					Type:      "TEXT",
					SectionId: "Test Section Id",
				},
			},
		}
		isValid := DidCreateWishlistTemplateDtoPassValidation(dto)
		if !isValid {
			t.Errorf("Expected %v to be true", isValid)
		}
	})

	t.Run("Should return false if the dto is invalid", func(t *testing.T) {
		dto := dto.CreateWishlistTemplateDto{
			Title:       "",
			Description: "Test Description",
			TemplateSections: []dto.CreateWishtlistTemplateSectionDto{
				{
					Title:     "Test Section Title",
					Type:      "TEXT",
					SectionId: "Test Section Id",
				},
			},
		}
		isValid := DidCreateWishlistTemplateDtoPassValidation(dto)
		if isValid {
			t.Errorf("Expected %v to be false", isValid)
		}
	})
}

func BenchmarkDidCreateWishlistTemplateDtoPassValidation_Valid(b *testing.B) {
	dto := dto.CreateWishlistTemplateDto{
		Title:       "Test Title",
		Description: "Test Description",
		TemplateSections: []dto.CreateWishtlistTemplateSectionDto{
			{
				Title:     "Test Section Title",
				Type:      "TEXT",
				SectionId: "Test Section Id",
			},
		},
	}
	for i := 0; i < b.N; i++ {
		DidCreateWishlistTemplateDtoPassValidation(dto)
	}
}
