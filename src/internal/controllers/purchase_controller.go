package controllers

import (
	"context"
	"sd/internal/domain/entities"
)

func (c *Controller) GetPurchase(ctx context.Context, id int) (*entities.Purchase, error) {
	return c.purchaseService.GetById(ctx, id)
}

func (c *Controller) AddPurchase(ctx context.Context, purchase *entities.Purchase) (int, error) {
	return c.purchaseService.Create(ctx, purchase)
}

func (c *Controller) GetAllPurchases(ctx context.Context) (entities.Purchases, error) {
	return c.purchaseService.GetAll(ctx)
}
