package services

import (
	"context"
	"sd/internal/domain/dto"
	"sd/internal/domain/entities"
)

type IAnimalRepo interface {
	GetById(ctx context.Context, id int) (*entities.Animal, error)
	GetAll(ctx context.Context) (entities.Animals, error)
	GetCrtrAll(ctx context.Context, crtr int) (entities.Animals, error)
	Create(ctx context.Context, dto *dto.CreateAnimal) error
	Update(ctx context.Context, id int, dto *dto.UpdateAnimal) error
	Delete(ctx context.Context, id int) error
}

type AnimalService struct {
	repo IAnimalRepo
}

func NewAnimalService(repo IAnimalRepo) *AnimalService {
	return &AnimalService{
		repo: repo,
	}
}

func (r *AnimalService) GetById(ctx context.Context, id int) (*entities.Animal, error) {
	return r.repo.GetById(ctx, id)
}
func (r *AnimalService) GetAll(ctx context.Context) (entities.Animals, error) {
	return r.repo.GetAll(ctx)
}
func (r *AnimalService) GetCrtrAll(ctx context.Context, crtr int) (entities.Animals, error) {
	return r.repo.GetCrtrAll(ctx, crtr)
}
func (r *AnimalService) Create(ctx context.Context, dto *dto.CreateAnimal) error {
	return r.repo.Create(ctx, dto)
}
func (r *AnimalService) Update(ctx context.Context, id int, dto *dto.UpdateAnimal) error {
	return r.repo.Update(ctx, id, dto)
}
func (r *AnimalService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
