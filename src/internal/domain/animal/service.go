package animal

import "context"

type AnimalService struct {
	repo IAnimalRepo
}

func NewAnimalService(repo IAnimalRepo) *AnimalService {
	return &AnimalService{
		repo: repo,
	}
}

func (r *AnimalService) GetById(ctx context.Context, id int32) (*Animal, error) {
	return r.repo.GetById(ctx, id)
}
func (r *AnimalService) GetAll(ctx context.Context) (Animals, error) {
	return r.repo.GetAll(ctx)
}
func (r *AnimalService) Create(ctx context.Context, dto *CreateAnimal) (*Animal, error) {
	return r.repo.Create(ctx, dto)
}
func (r *AnimalService) Update(ctx context.Context, id int32, dto *UpdateAnimal) (*Animal, error) {
	return r.repo.Update(ctx, id, dto)
}
func (r *AnimalService) Delete(ctx context.Context, id int32) error {
	return r.repo.Delete(ctx, id)
}
