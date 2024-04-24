package repository

import "github.com/Figaarillo/sacachispa-soft/internal/entity"

type ProductRepository interface {
	List(offset, limit int) ([]entity.Product, error)
	GetByID(id entity.ID) (entity.Product, error)
	Create(payload *entity.Product) error
	Update(id entity.ID, payload entity.Product) error
	Delete(id entity.ID) error
}
