package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"testing"
	"time"
)

func TestPurchaseService(t *testing.T) {
	ctx := context.Background()
	err := TruncateTables(client, ctx)
	if err != nil {
		return
	}

	newShelter := &entities.Shelter{
		Street: "Первомайская",
		House:  10,
	}
	shelId, err := shelterService.Create(ctx, newShelter)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	newAnimal := &entities.Animal{
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: shelId,
		Type:      "собака",
		Gender:    "мужской",
	}
	anId, err := animalService.Create(ctx, newAnimal)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	lastTime, _ := time.Parse("2023-07-25", "2006-01-02")
	newPurchase := &entities.Purchase{
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  lastTime,
		AnimalId:  anId,
	}
	var id int
	id, err = purchaseService.Create(ctx, newPurchase)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	newPurchase.Id = id
	t1, err := purchaseService.GetById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, newPurchase, t1)

	ts1, err := purchaseService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, newPurchase, ts1[0])

	newPurchase.Cost += 30
	err = purchaseService.Update(ctx, newPurchase)
	assert.NoError(t, err)

	err = purchaseService.Delete(ctx, id)
	assert.NoError(t, err)

	_, err = purchaseService.GetById(ctx, id)
	assert.Error(t, err)

	ts2, err := purchaseService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)
}
