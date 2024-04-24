package entity

import (
	"time"

	"github.com/Figaarillo/sacachispa-soft/internal/util"
	"gorm.io/gorm"
)

type Product struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	Brand       Brand          `json:"brand" gorm:"foreignKey:BrandID"`
	Stock       int            `json:"stock"`
	Price       float64        `json:"price"`
	BrandID     ID             `json:"brand_id" gorm:"not null"`
	ID          ID             `json:"id"`
}

func NewProduct(payload Product) (*Product, error) {
	product := &Product{
		ID:          NewID(),
		Name:        payload.Name,
		Description: payload.Description,
		Stock:       payload.Stock,
		Price:       payload.Price,
		BrandID:     payload.BrandID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Update(payload Product) error {
	util.AssignIfNotEmpty(&p.Name, payload.Name)
	util.AssignIfNotEmpty(&p.Description, payload.Description)
	util.AssignIfNotZero(&p.Stock, payload.Stock)
	util.AssignIfNotZeroFloat(&p.Price, payload.Price)
	p.UpdatedAt = time.Now()

	if err := p.Validate(); err != nil {
		return err
	}

	return nil
}

func (p *Product) Validate() error {
	if p.Name == "" || p.Description == "" || p.BrandID.String() == "" || p.Stock == 0 || p.Price == 0 {
		return ErrMissingField
	}

	return nil
}
