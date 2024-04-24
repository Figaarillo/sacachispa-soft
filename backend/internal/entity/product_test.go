package entity_test

import (
	"testing"

	"github.com/Figaarillo/sacachispa-soft/internal/entity"
)

func TestProduct(t *testing.T) {
	tName := "product"
	tDesc := "product description"
	tPrice := 10.0
	tStock := 5

	tNewName := "updated product"
	tNewDesc := "updated product description"
	tNewPrice := 20.0
	tNewStock := 10

	t.Run("Test NewProduct() with valid fields", func(t *testing.T) {
		brand, _ := entity.NewBrand(entity.Brand{Name: "brand", Description: "brand description"})

		product, err := entity.NewProduct(entity.Product{
			Name:        tName,
			Description: tDesc,
			Price:       tPrice,
			Stock:       tStock,
			BrandID:     brand.ID,
		})
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		if product == nil {
			t.Fatal("Failed to create product: product is nil")
		}

		if product.Name != tName {
			t.Errorf("Expected product name '%s', got %s", tName, product.Name)
		}

		if product.Description != tDesc {
			t.Errorf("Expected product description '%s', got %s", tDesc, product.Description)
		}

		if product.Price != tPrice {
			t.Errorf("Expected product price '%f', got %f", tPrice, product.Price)
		}

		if product.Stock != tStock {
			t.Errorf("Expected product stock '%d', got %d", tStock, product.Stock)
		}

		if product.BrandID != brand.ID {
			t.Errorf("Expected product brand ID '%s', got %s", brand.ID, product.BrandID)
		}
	})

	t.Run("Test NewProduct() with invalid fields", func(t *testing.T) {
		_, err := entity.NewProduct(entity.Product{
			Name:        "",
			Description: "",
			Price:       0,
			Stock:       0,
			BrandID:     "",
		})
		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}

		if err != entity.ErrMissingField {
			t.Errorf("Expected error %v, got %v", entity.ErrMissingField, err)
		}
	})

	t.Run("Test Update() with valid fields", func(t *testing.T) {
		brand, _ := entity.NewBrand(entity.Brand{Name: "brand", Description: "brand Description"})

		product, _ := entity.NewProduct(entity.Product{
			Name:        tName,
			Description: tDesc,
			Price:       tPrice,
			Stock:       tStock,
			BrandID:     brand.ID,
		})

		err := product.Update(entity.Product{
			Name:        tNewName,
			Description: tNewDesc,
			Price:       tNewPrice,
			Stock:       tNewStock,
		})
		if err != nil {
			t.Fatalf("Failed to update product: %v", err)
		}

		if product.Name != tNewName {
			t.Errorf("Expected product name '%s', got %s", tNewName, product.Name)
		}

		if product.Description != tNewDesc {
			t.Errorf("Expected product description '%s', got %s", tNewDesc, product.Description)
		}

		if product.Price != tNewPrice {
			t.Errorf("Expected product price '%f', got %f", tNewPrice, product.Price)
		}

		if product.Stock != tNewStock {
			t.Errorf("Expected product stock '%d', got %d", tNewStock, product.Stock)
		}
	})
}
