package helper

import "testing"

func TestDidUserDisplayNamePassValidation(t *testing.T) {
	t.Run("Should return true if the display name is valid", func(t *testing.T) {
		displayName := "John Doe"
		isValid := DidUserDisplayNamePassValidation(displayName)
		if !isValid {
			t.Errorf("Expected %v to be true", isValid)
		}
	})

	t.Run("Should return false if the display name is invalid", func(t *testing.T) {
		displayName := "John Doe!"
		isValid := DidUserDisplayNamePassValidation(displayName)
		if isValid {
			t.Errorf("Expected %v to be false", isValid)
		}
	})

	t.Run("Should return false if the display name is empty", func(t *testing.T) {
		displayName := ""
		isValid := DidUserDisplayNamePassValidation(displayName)
		if isValid {
			t.Errorf("Expected %v to be false", isValid)
		}
	})

	t.Run("Should return false if the display name matches a bad option", func(t *testing.T) {
		displayName := "administrator"
		isValid := DidUserDisplayNamePassValidation(displayName)
		if isValid {
			t.Errorf("Expected %v to be false", isValid)
		}
	})
}

func BenchmarkDidUserDisplayNamePassValidation_Valid(b *testing.B) {
	displayName := "John Doe"
	for i := 0; i < b.N; i++ {
		DidUserDisplayNamePassValidation(displayName)
	}
}

func BenchmarkDidUserDisplayNamePassValidation_Invalid(b *testing.B) {
	displayName := "John Doe!"
	for i := 0; i < b.N; i++ {
		DidUserDisplayNamePassValidation(displayName)
	}
}

func BenchmarkDidUserDisplayNamePassValidation_Empty(b *testing.B) {
	displayName := ""
	for i := 0; i < b.N; i++ {
		DidUserDisplayNamePassValidation(displayName)
	}
}

func BenchmarkDidUserDisplayNamePassValidation_BadOption(b *testing.B) {
	displayName := "administrator"
	for i := 0; i < b.N; i++ {
		DidUserDisplayNamePassValidation(displayName)
	}
}
