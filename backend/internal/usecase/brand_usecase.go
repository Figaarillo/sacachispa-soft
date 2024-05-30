package usecase

import (
	"github.com/Figaarillo/sacachispa-soft/internal/entity"
	"github.com/Figaarillo/sacachispa-soft/internal/exeption"
	"github.com/Figaarillo/sacachispa-soft/internal/repository"
)

type BrandUsecase struct {
	repository repository.BrandRepository
}

func NewBrandUsecase(repository repository.BrandRepository) *BrandUsecase {
	return &BrandUsecase{repository: repository}
}

func (uc *BrandUsecase) List(offset, limit int) ([]entity.Brand, error) {
	if offset < 0 || limit < 0 || (offset == 0 && limit == 0) {
		return nil, exeption.ErrInvalidPagination
	}

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

	if err := uc.repository.Create(brand); err != nil {
		return err
	}

	return nil
}

func (uc *BrandUsecase) Update(id string, payload entity.Brand) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	if err := uc.repository.Update(idParsed, payload); err != nil {
		return err
	}

	return nil
}

func (uc *BrandUsecase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	if err := uc.repository.Delete(idParsed); err != nil {
		return err
	}

	return nil
}
