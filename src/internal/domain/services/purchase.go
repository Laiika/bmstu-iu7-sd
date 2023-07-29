package services

import (
	"context"
	"sd/internal/domain/entities"
)

// mockgen -destination mocks/purchase_mock.go -package mocks . IPurchaseRepo

type IPurchaseRepo interface {
	GetById(ctx context.Context, id int) (*entities.Purchase, error)
	GetAll(ctx context.Context) (entities.Purchases, error)
	Create(ctx context.Context, purchase *entities.Purchase) (int, error)
	Update(ctx context.Context, purchase *entities.Purchase) error
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
func (r *PurchaseService) Create(ctx context.Context, purchase *entities.Purchase) (int, error) {
	if err := purchase.IsValid(); err != nil {
		return 0, err
	}

	return r.repo.Create(ctx, purchase)
}
func (r *PurchaseService) Update(ctx context.Context, purchase *entities.Purchase) error {
	if err := purchase.IsValid(); err != nil {
		return err
	}

	return r.repo.Update(ctx, purchase)
}
func (r *PurchaseService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
