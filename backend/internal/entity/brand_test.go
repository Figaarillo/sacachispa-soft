package entity_test

import (
	"testing"

	"github.com/Figaarillo/sacachispa-soft/internal/entity"
)

func TestBrand(t *testing.T) {
	tName := "brand"
	tDesc := "brand description"
	tNewName := "brand 2"
	tNewDesc := "brand description 2"

	t.Run("Test NewBrand() with valid fields", func(t *testing.T) {
		brand, err := entity.NewBrand(entity.Brand{Name: tName, Description: tDesc})
		if err != nil {
			t.Fatalf("Failed to create brand: %v", err)
		}

		if brand == nil {
			t.Fatal("Failed to create brand: brand is nil")
		}

		if brand.Name != "brand" {
			t.Errorf("Expected brand name '%s', got '%s'", tName, brand.Name)
		}

		if brand.Description != "brand description" {
			t.Errorf("Expected brand description '%s', got %s", tDesc, brand.Description)
		}
	})

	t.Run("Test NewBrand() with invalid fields", func(t *testing.T) {
		_, err := entity.NewBrand(entity.Brand{Name: "", Description: ""})
		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}

		if err != entity.ErrMissingField {
			t.Errorf("Expected error '%v', got '%v'", entity.ErrMissingField, err)
		}
	})

	t.Run("Test Update() with valid fields", func(t *testing.T) {
		brand, err := entity.NewBrand(entity.Brand{Name: tName, Description: tDesc})
		if err != nil {
			t.Fatalf("Failed to create brand: %v", err)
		}

		if err := brand.Update(entity.Brand{Name: tNewName, Description: tNewDesc}); err != nil {
			t.Errorf("Failed to update brand: %v", err)
		}

		if brand.Name != tNewName {
			t.Errorf("Expected brand name '%s', got '%s'", tNewName, brand.Name)
		}
	})
}
