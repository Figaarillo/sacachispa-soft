package repository

import "github.com/Figaarillo/sacachispa-soft/internal/entity"

type BrandRepository interface {
	List(offset, limit int) ([]entity.Brand, error)
	GetByID(id entity.ID) (entity.Brand, error)
	GetByName(name string) (entity.Brand, error)
	Create(payload *entity.Brand) error
	Update(id entity.ID, payload entity.Brand) error
	Delete(id string) error
}
