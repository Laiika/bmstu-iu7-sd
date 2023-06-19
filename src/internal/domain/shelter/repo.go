package shelter

import "context"

type IShelterRepo interface {
	GetById(ctx context.Context, id int) (*Shelter, error)
	GetAll(ctx context.Context) (Shelters, error)
	Create(ctx context.Context, dto *CreateShelter) (*Shelter, error)
	Update(ctx context.Context, id int, dto *UpdateShelter) (*Shelter, error)
	Delete(ctx context.Context, id int) error
}
