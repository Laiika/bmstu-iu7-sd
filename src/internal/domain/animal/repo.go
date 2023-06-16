package animal

import "context"

type IAnimalRepo interface {
	GetById(ctx context.Context, id int32) (*Animal, error)
	List(ctx context.Context) (Animals, error)
	Create(ctx context.Context, dto *CreateAnimal) (*Animal, error)
	Update(ctx context.Context, num int32, dto *UpdateAnimal) (*Animal, error)
	Delete(ctx context.Context, num int32) error
}
