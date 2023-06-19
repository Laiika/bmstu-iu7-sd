package animal

import "context"

type IAnimalRepo interface {
	GetById(ctx context.Context, id int) (*Animal, error)
	GetAll(ctx context.Context) (Animals, error)
	GetCrtrAll(ctx context.Context, crtr int) (Animals, error)
	Create(ctx context.Context, dto *CreateAnimal) (*Animal, error)
	Update(ctx context.Context, id int, dto *UpdateAnimal) (*Animal, error)
	Delete(ctx context.Context, id int) error
}
