package controllers

import "sd/internal/domain/services"

type Controller struct {
	curatorService  *services.CuratorService
	animalService   *services.AnimalService
	shelterService  *services.ShelterService
	purchaseService *services.PurchaseService
}

func NewController(
	curatorService *services.CuratorService,
	animalService *services.AnimalService,
	shelterService *services.ShelterService,
	purchaseService *services.PurchaseService,
) *Controller {

	return &Controller{
		curatorService:  curatorService,
		animalService:   animalService,
		shelterService:  shelterService,
		purchaseService: purchaseService,
	}
}
