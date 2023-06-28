package services

import (
	"context"
	"sd/internal/domain/dto"
	"sd/internal/domain/entities"
)

type IPurchaseRepo interface {
	GetById(ctx context.Context, id int) (*entities.Purchase, error)
	GetAll(ctx context.Context) (entities.Purchases, error)
	Create(ctx context.Context, dto *dto.CreatePurchase) error
	Update(ctx context.Context, id int, dto *dto.UpdatePurchase) error
	Delete(ctx context.Context, id int) error
}

type PurchaseService struct {
	repo IPurchaseRepo
}

func NewPurchaseService(repo IPurchaseRepo) *PurchaseService {
	return &PurchaseService{
		repo: repo,
	}
}

func (r *PurchaseService) GetById(ctx context.Context, id int) (*entities.Purchase, error) {
	return r.repo.GetById(ctx, id)
}
func (r *PurchaseService) GetAll(ctx context.Context) (entities.Purchases, error) {
	return r.repo.GetAll(ctx)
}
func (r *PurchaseService) Create(ctx context.Context, dto *dto.CreatePurchase) error {
	return r.repo.Create(ctx, dto)
}
func (r *PurchaseService) Update(ctx context.Context, id int, dto *dto.UpdatePurchase) error {
	return r.repo.Update(ctx, id, dto)
}
func (r *PurchaseService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
