package usecase

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repository: repository}
}

func (uc *ProductUsecase) List(offset, limit int) ([]entity.Product, error) {
	return uc.repository.List(offset, limit)
}

func (uc *ProductUsecase) GetByID(id string) (entity.Product, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.Product{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *ProductUsecase) Create(payload entity.Product) error {
	product, err := entity.NewProduct(payload)
	if err != nil {
		return err
	}

	uc.repository.Create(product)

	return nil
}

func (uc *ProductUsecase) Update(id string, payload entity.Product) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	uc.repository.Update(idParsed, payload)

	return nil
}

func (uc *ProductUsecase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	uc.repository.Delete(idParsed)

	return nil
}
