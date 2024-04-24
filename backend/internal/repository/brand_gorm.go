package repository

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"gorm.io/gorm"
)

type BrandGorm struct {
	db *gorm.DB
}

func NewBrandGorm(db *gorm.DB) *BrandGorm {
	return &BrandGorm{db: db}
}

func (b *BrandGorm) List(offset, limit int) ([]entity.Brand, error) {
	var brands []entity.Brand

	if result := b.db.Offset(offset).Limit(limit).
		Find(&brands); result.Error != nil {
		return nil, result.Error
	}

	return brands, nil
}

func (b *BrandGorm) GetByID(id entity.ID) (entity.Brand, error) {
	var brand entity.Brand

	if result := b.db.First(&brand, "id = ?", id); result.Error != nil {
		return entity.Brand{}, result.Error
	}

	return brand, nil
}

func (b *BrandGorm) GetByName(name string) (entity.Brand, error) {
	var brand entity.Brand

	if result := b.db.First(&brand, "name = ?", name); result.Error != nil {
		return entity.Brand{}, result.Error
	}

	return brand, nil
}

func (b *BrandGorm) Create(payload *entity.Brand) error {
	if result := b.db.Create(payload); result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BrandGorm) Update(id entity.ID, payload entity.Brand) error {
	var brand entity.Brand

	if result := b.db.First(&brand, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if err := brand.Update(payload); err != nil {
		return err
	}

	if result := b.db.Save(&brand); result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BrandGorm) Delete(id entity.ID) error {
	var brand entity.Brand

	if result := b.db.First(&brand, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if result := b.db.Delete(&brand); result.Error != nil {
		return result.Error
	}

	return nil
}
