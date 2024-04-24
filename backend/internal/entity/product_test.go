package entity_test

import (
	"testing"

	"github.com/Figaarillo/sacachispa-soft/internal/entity"
)

func TestProduct(t *testing.T) {
	testName := "product"
	testDescription := "product description"
	testPrice := 10.0
	testStock := 5

	testUpdatedName := "updated product"
	testUpdatedDescription := "updated product description"
	testUpdatedPrice := 20.0
	testUpdatedStock := 10

	t.Run("Test NewProduct() with valid fields", func(t *testing.T) {
		brand, _ := entity.NewBrand(entity.Brand{Name: "brand", Description: "brand description"})

		product, err := entity.NewProduct(entity.Product{
			Name:        testName,
			Description: testDescription,
			Price:       testPrice,
			Stock:       testStock,
			BrandID:     brand.ID,
		})
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		if product == nil {
			t.Fatal("Failed to create product: product is nil")
		}

		if product.Name != testName {
			t.Errorf("Expected product name '%s', got %s", testName, product.Name)
		}

		if product.Description != testDescription {
			t.Errorf("Expected product description '%s', got %s", testDescription, product.Description)
		}

		if product.Price != testPrice {
			t.Errorf("Expected product price '%f', got %f", testPrice, product.Price)
		}

		if product.Stock != testStock {
			t.Errorf("Expected product stock '%d', got %d", testStock, product.Stock)
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
			Name:        testName,
			Description: testDescription,
			Price:       testPrice,
			Stock:       testStock,
			BrandID:     brand.ID,
		})

		err := product.Update(entity.Product{
			Name:        testUpdatedName,
			Description: testUpdatedDescription,
			Price:       testUpdatedPrice,
			Stock:       testUpdatedStock,
		})
		if err != nil {
			t.Fatalf("Failed to update product: %v", err)
		}

		if product.Name != testUpdatedName {
			t.Errorf("Expected product name '%s', got %s", testUpdatedName, product.Name)
		}

		if product.Description != testUpdatedDescription {
			t.Errorf("Expected product description '%s', got %s", testUpdatedDescription, product.Description)
		}

		if product.Price != testUpdatedPrice {
			t.Errorf("Expected product price '%f', got %f", testUpdatedPrice, product.Price)
		}

		if product.Stock != testUpdatedStock {
			t.Errorf("Expected product stock '%d', got %d", testUpdatedStock, product.Stock)
		}
	})
}
