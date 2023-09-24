package controllers

import "sd/internal/domain/services"

type Controller struct {
	curatorService  *services.CuratorService
	animalService   *services.AnimalService
	shelterService  *services.ShelterService
	purchaseService *services.PurchaseService
	diseaseService  *services.DiseaseService
}

func NewController(
	curatorService *services.CuratorService,
	animalService *services.AnimalService,
	shelterService *services.ShelterService,
	purchaseService *services.PurchaseService,
	diseaseService *services.DiseaseService,
) *Controller {

	return &Controller{
		curatorService:  curatorService,
		animalService:   animalService,
		shelterService:  shelterService,
		purchaseService: purchaseService,
		diseaseService:  diseaseService,
	}
}
