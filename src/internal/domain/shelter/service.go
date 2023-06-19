package shelter

import "context"

type ShelterService struct {
	repo IShelterRepo
}

func NewShelterService(repo IShelterRepo) *ShelterService {
	return &ShelterService{
		repo: repo,
	}
}

func (r *ShelterService) GetById(ctx context.Context, id int) (*Shelter, error) {
	return r.repo.GetById(ctx, id)
}
func (r *ShelterService) GetAll(ctx context.Context) (Shelters, error) {
	return r.repo.GetAll(ctx)
}
func (r *ShelterService) Create(ctx context.Context, dto *CreateShelter) (*Shelter, error) {
	return r.repo.Create(ctx, dto)
}
func (r *ShelterService) Update(ctx context.Context, id int, dto *UpdateShelter) (*Shelter, error) {
	return r.repo.Update(ctx, id, dto)
}
func (r *ShelterService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
