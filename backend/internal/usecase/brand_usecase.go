package usecase

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
)

type BrandUsecase struct {
	repository repository.BrandRepository
}

func NewBrandUsecase(repository repository.BrandRepository) *BrandUsecase {
	return &BrandUsecase{repository: repository}
}

func (uc *BrandUsecase) List(offset, limit int) ([]entity.Brand, error) {
	return uc.repository.List(offset, limit)
}

func (uc *BrandUsecase) GetByID(id string) (entity.Brand, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.Brand{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *BrandUsecase) GetByName(name string) (entity.Brand, error) {
	return uc.repository.GetByName(name)
}

func (uc *BrandUsecase) Create(payload entity.Brand) error {
	brand, err := entity.NewBrand(payload)
	if err != nil {
		return err
	}

	uc.repository.Create(brand)

	return nil
}

func (uc *BrandUsecase) Update(id string, payload entity.Brand) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	uc.repository.Update(idParsed, payload)

	return nil
}

func (uc *BrandUsecase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	uc.repository.Delete(idParsed)

	return nil
}
