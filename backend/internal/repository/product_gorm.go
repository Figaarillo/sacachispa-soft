package repository

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"gorm.io/gorm"
)

type ProductGorm struct {
	db *gorm.DB
}

func NewProductGorm(db *gorm.DB) *ProductGorm {
	return &ProductGorm{db: db}
}

func (p *ProductGorm) List(offset, limit int) ([]entity.Product, error) {
	var products []entity.Product

	if result := p.db.Preload("Brand").Offset(offset).Limit(limit).
		Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p *ProductGorm) GetByID(id entity.ID) (entity.Product, error) {
	var product entity.Product

	if result := p.db.Preload("Brand").First(&product, "id = ?", id); result.Error != nil {
		return entity.Product{}, result.Error
	}

	return product, nil
}

func (p *ProductGorm) Create(payload *entity.Product) error {
	if result := p.db.Create(payload); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductGorm) Update(id entity.ID, payload entity.Product) error {
	var product entity.Product

	if result := p.db.First(&product, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if err := product.Update(payload); err != nil {
		return err
	}

	if result := p.db.Save(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductGorm) Delete(id entity.ID) error {
	var product entity.Product

	if result := p.db.First(&product, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if result := p.db.Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}
