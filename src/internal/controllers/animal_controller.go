package controllers

import (
	"context"
	"sd/internal/domain/entities"
)

func (c *Controller) GetAnimal(ctx context.Context, id int) (*entities.Animal, error) {
	return c.animalService.GetById(ctx, id)
}

func (c *Controller) AddAnimal(ctx context.Context, animal *entities.Animal) (int, error) {
	return c.animalService.Create(ctx, animal)
}

func (c *Controller) GetAllAnimals(ctx context.Context) (entities.Animals, error) {
	return c.animalService.GetAll(ctx)
}

func (c *Controller) GetAllCrtrAnimals(ctx context.Context, id int) (entities.Animals, error) {
	return c.animalService.GetCrtrAll(ctx, id)
}
