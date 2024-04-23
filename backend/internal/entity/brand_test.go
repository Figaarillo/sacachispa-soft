package entity_test

import (
	"testing"

	"github.com/Figaarillo/sacachispa-soft/internal/entity"
)

func TestBrand(t *testing.T) {
	t.Run("Test NewBrand() with valid fields", func(t *testing.T) {
		brand, err := entity.NewBrand(entity.Brand{Name: "brand", Description: "brand description"})
		if err != nil {
			t.Fatalf("Failed to create brand: %v", err)
		}

		if brand == nil {
			t.Fatal("Failed to create brand: brand is nil")
		}

		if brand.Name != "brand" {
			t.Errorf("Expected brand name 'brand', got %s", brand.Name)
		}

		if brand.Description != "brand description" {
			t.Errorf("Expected brand description 'brand description', got %s", brand.Description)
		}
	})

	t.Run("Test NewBrand() with invalid fields", func(t *testing.T) {
		_, err := entity.NewBrand(entity.Brand{Name: "", Description: ""})
		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}
	})

	t.Run("Test Update() with valid fields", func(t *testing.T) {
		brand, err := entity.NewBrand(entity.Brand{Name: "brand", Description: "brand description"})
		if err != nil {
			t.Fatalf("Failed to create brand: %v", err)
		}

		newName := "brand 2"
		newDescription := "brand description 2"

		if err := brand.Update(entity.Brand{Name: newName, Description: newDescription}); err != nil {
			t.Errorf("Failed to update brand: %v", err)
		}

		if brand.Name != newName {
			t.Errorf("Expected brand name '%s', got '%s'", newName, brand.Name)
		}
	})
}
