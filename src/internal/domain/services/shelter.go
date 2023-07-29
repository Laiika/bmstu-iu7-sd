package services

import (
	"context"
	"sd/internal/domain/entities"
)

// mockgen -destination mocks/shelter_mock.go -package mocks . IShelterRepo

type IShelterRepo interface {
	GetById(ctx context.Context, id int) (*entities.Shelter, error)
	GetAll(ctx context.Context) (entities.Shelters, error)
	Create(ctx context.Context, shelter *entities.Shelter) (int, error)
	Update(ctx context.Context, shelter *entities.Shelter) error
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
func (r *ShelterService) Create(ctx context.Context, shelter *entities.Shelter) (int, error) {
	if err := shelter.IsValid(); err != nil {
		return 0, err
	}

	return r.repo.Create(ctx, shelter)
}
func (r *ShelterService) Update(ctx context.Context, shelter *entities.Shelter) error {
	if err := shelter.IsValid(); err != nil {
		return err
	}

	return r.repo.Update(ctx, shelter)
}
func (r *ShelterService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
