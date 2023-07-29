package services

import (
	"context"
	"sd/internal/domain/entities"
)

// mockgen -destination mocks/curator_mock.go -package mocks . ICuratorRepo

type ICuratorRepo interface {
	GetById(ctx context.Context, id int) (*entities.Curator, error)
	GetByChatId(ctx context.Context, chatId string) (*entities.Curator, error)
	GetAll(ctx context.Context) (entities.Curators, error)
	Create(ctx context.Context, curator *entities.Curator) (int, error)
	Update(ctx context.Context, curator *entities.Curator) error
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
func (r *CuratorService) GetByChatId(ctx context.Context, chatId string) (*entities.Curator, error) {
	return r.repo.GetByChatId(ctx, chatId)
}
func (r *CuratorService) GetAll(ctx context.Context) (entities.Curators, error) {
	return r.repo.GetAll(ctx)
}
func (r *CuratorService) Create(ctx context.Context, curator *entities.Curator) (int, error) {
	if err := curator.IsValid(); err != nil {
		return 0, err
	}

	return r.repo.Create(ctx, curator)
}
func (r *CuratorService) Update(ctx context.Context, curator *entities.Curator) error {
	if err := curator.IsValid(); err != nil {
		return err
	}

	return r.repo.Update(ctx, curator)
}
func (r *CuratorService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
