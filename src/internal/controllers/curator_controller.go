package controllers

import (
	"context"
	"sd/internal/domain/entities"
)

func (c *Controller) GetCurator(ctx context.Context, id int) (*entities.Curator, error) {
	return c.curatorService.GetById(ctx, id)
}

func (c *Controller) GetCuratorByChatId(ctx context.Context, chatId string) (*entities.Curator, error) {
	return c.curatorService.GetByChatId(ctx, chatId)
}

func (c *Controller) AddCurator(ctx context.Context, curator *entities.Curator) (int, error) {
	return c.curatorService.Create(ctx, curator)
}
