package services

import (
	"context"
	"sd/internal/domain/dto"
	"sd/internal/domain/entities"
)

type IShelterRepo interface {
	GetById(ctx context.Context, id int) (*entities.Shelter, error)
	GetAll(ctx context.Context) (entities.Shelters, error)
	Create(ctx context.Context, dto *dto.CreateShelter) error
	Update(ctx context.Context, id int, dto *dto.UpdateShelter) error
	Delete(ctx context.Context, id int) error
}

type ShelterService struct {
	repo IShelterRepo
}

func NewShelterService(repo IShelterRepo) *ShelterService {
	return &ShelterService{
		repo: repo,
	}
}

func (r *ShelterService) GetById(ctx context.Context, id int) (*entities.Shelter, error) {
	return r.repo.GetById(ctx, id)
}
func (r *ShelterService) GetAll(ctx context.Context) (entities.Shelters, error) {
	return r.repo.GetAll(ctx)
}
func (r *ShelterService) Create(ctx context.Context, dto *dto.CreateShelter) error {
	return r.repo.Create(ctx, dto)
}
func (r *ShelterService) Update(ctx context.Context, id int, dto *dto.UpdateShelter) error {
	return r.repo.Update(ctx, id, dto)
}
func (r *ShelterService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
