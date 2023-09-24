package services

import (
	"context"
	"sd/internal/domain/entities"
)

// mockgen -destination mocks/disease_mock.go -package mocks . IDiseaseRepo

type IDiseaseRepo interface {
	GetById(ctx context.Context, id int) (*entities.Disease, error)
	GetAll(ctx context.Context) (entities.Diseases, error)
	GetAnimalAll(ctx context.Context, an int) (entities.Diseases, error)
	Create(ctx context.Context, disease *entities.Disease) (int, error)
	Update(ctx context.Context, disease *entities.Disease) error
	Delete(ctx context.Context, id int) error
}

type DiseaseService struct {
	repo IDiseaseRepo
}

func NewDiseaseService(repo IDiseaseRepo) *DiseaseService {
	return &DiseaseService{
		repo: repo,
	}
}

func (r *DiseaseService) GetById(ctx context.Context, id int) (*entities.Disease, error) {
	return r.repo.GetById(ctx, id)
}

func (r *DiseaseService) GetAll(ctx context.Context) (entities.Diseases, error) {
	return r.repo.GetAll(ctx)
}

func (r *DiseaseService) GetAnimalAll(ctx context.Context, an int) (entities.Diseases, error) {
	return r.repo.GetAnimalAll(ctx, an)
}

func (r *DiseaseService) Create(ctx context.Context, disease *entities.Disease) (int, error) {
	if err := disease.IsValid(); err != nil {
		return 0, err
	}

	return r.repo.Create(ctx, disease)
}

func (r *DiseaseService) Update(ctx context.Context, disease *entities.Disease) error {
	if err := disease.IsValid(); err != nil {
		return err
	}

	return r.repo.Update(ctx, disease)
}

func (r *DiseaseService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
