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

func (r *AnimalService) GetById(ctx context.Context, id int) (*Animal, error) {
	return r.repo.GetById(ctx, id)
}
func (r *AnimalService) GetAll(ctx context.Context) (Animals, error) {
	return r.repo.GetAll(ctx)
}
func (r *AnimalService) GetCrtrAll(ctx context.Context, crtr int) (Animals, error) {
	return r.repo.GetCrtrAll(ctx, crtr)
}
func (r *AnimalService) Create(ctx context.Context, dto *CreateAnimal) (*Animal, error) {
	return r.repo.Create(ctx, dto)
}
func (r *AnimalService) Update(ctx context.Context, id int, dto *UpdateAnimal) (*Animal, error) {
	return r.repo.Update(ctx, id, dto)
}
func (r *AnimalService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
