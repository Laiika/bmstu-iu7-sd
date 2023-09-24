package controllers

import (
	"context"
	"sd/internal/domain/entities"
)

func (c *Controller) AddDisease(ctx context.Context, disease *entities.Disease) (int, error) {
	return c.diseaseService.Create(ctx, disease)
}

func (c *Controller) GetAllAnimalDiseases(ctx context.Context, id int) (entities.Diseases, error) {
	return c.diseaseService.GetAnimalAll(ctx, id)
}
