package curator

import "context"

type CuratorService struct {
	repo ICuratorRepo
}

func NewCuratorService(repo ICuratorRepo) *CuratorService {
	return &CuratorService{
		repo: repo,
	}
}

func (r *CuratorService) GetById(ctx context.Context, id int) (*Curator, error) {
	return r.repo.GetById(ctx, id)
}
func (r *CuratorService) GetAll(ctx context.Context) (Curators, error) {
	return r.repo.GetAll(ctx)
}
func (r *CuratorService) Create(ctx context.Context, dto *CreateCurator) (*Curator, error) {
	return r.repo.Create(ctx, dto)
}
func (r *CuratorService) Update(ctx context.Context, id int, dto *UpdateCurator) (*Curator, error) {
	return r.repo.Update(ctx, id, dto)
}
func (r *CuratorService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
