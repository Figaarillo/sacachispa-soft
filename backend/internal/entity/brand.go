package entity

import "time"

type Brand struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ID          ID        `json:"id"`
}

func NewBrand(payload Brand) (*Brand, error) {
	brand := &Brand{
		ID:          NewID(),
		Name:        payload.Name,
		Description: payload.Description,
	}

	if err := brand.Validate(); err != nil {
		return nil, err
	}

	return brand, nil
}

func (b *Brand) Update(payload Brand) error {
	b.Name = payload.Name
	b.Description = payload.Description

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
