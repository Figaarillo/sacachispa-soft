package entity

import (
	"time"

	"github.com/Figaarillo/sacachispa-soft/internal/util"
	"gorm.io/gorm"
)

type Brand struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Products    []Product      `json:"products" gorm:"foreignKey:BrandID"`
	ID          ID             `json:"id"`
}

func NewBrand(payload Brand) (*Brand, error) {
	brand := &Brand{
		ID:          NewID(),
		Name:        payload.Name,
		Description: payload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := brand.Validate(); err != nil {
		return nil, err
	}

	return brand, nil
}

func (b *Brand) Update(payload Brand) error {
	util.AssignIfNotEmpty(&b.Name, payload.Name)
	util.AssignIfNotEmpty(&b.Description, payload.Description)
	b.UpdatedAt = time.Now()

	if err := b.Validate(); err != nil {
		return err
	}

	return nil
}

func (b *Brand) Validate() error {
	if b.Name == "" || b.Description == "" {
		return ErrMissingField
	}
	return nil
}
