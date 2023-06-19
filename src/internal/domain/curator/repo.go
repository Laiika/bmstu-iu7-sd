package curator

import "context"

type ICuratorRepo interface {
	GetById(ctx context.Context, id int) (*Curator, error)
	GetAll(ctx context.Context) (Curators, error)
	Create(ctx context.Context, dto *CreateCurator) (*Curator, error)
	Update(ctx context.Context, id int, dto *UpdateCurator) (*Curator, error)
	Delete(ctx context.Context, id int) error
}
