package services

import (
	"context"
	"sd/internal/domain/dto"
	"sd/internal/domain/entities"
)

type ICuratorRepo interface {
	GetById(ctx context.Context, id int) (*entities.Curator, error)
	GetAll(ctx context.Context) (entities.Curators, error)
	Create(ctx context.Context, dto *dto.CreateCurator) error
	Update(ctx context.Context, id int, dto *dto.UpdateCurator) error
	Delete(ctx context.Context, id int) error
}

type CuratorService struct {
	repo ICuratorRepo
}

func NewCuratorService(repo ICuratorRepo) *CuratorService {
	return &CuratorService{
		repo: repo,
	}
}

func (r *CuratorService) GetById(ctx context.Context, id int) (*entities.Curator, error) {
	return r.repo.GetById(ctx, id)
}
func (r *CuratorService) GetAll(ctx context.Context) (entities.Curators, error) {
	return r.repo.GetAll(ctx)
}
func (r *CuratorService) Create(ctx context.Context, dto *dto.CreateCurator) error {
	return r.repo.Create(ctx, dto)
}
func (r *CuratorService) Update(ctx context.Context, id int, dto *dto.UpdateCurator) error {
	return r.repo.Update(ctx, id, dto)
}
func (r *CuratorService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
