package controllers

import (
	"context"
	"sd/internal/domain/entities"
)

func (c *Controller) GetShelter(ctx context.Context, id int) (*entities.Shelter, error) {
	return c.shelterService.GetById(ctx, id)
}

func (c *Controller) AddShelter(ctx context.Context, shelter *entities.Shelter) (int, error) {
	return c.shelterService.Create(ctx, shelter)
}

func (c *Controller) GetAllShelters(ctx context.Context) (entities.Shelters, error) {
	return c.shelterService.GetAll(ctx)
}
