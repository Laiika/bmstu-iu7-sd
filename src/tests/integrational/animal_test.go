package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"testing"
)

func TestAnimalService(t *testing.T) {
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
	var id int
	id, err = animalService.Create(ctx, newAnimal)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	newAnimal.Id = id
	t1, err := animalService.GetById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, newAnimal, t1)

	ts1, err := animalService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, newAnimal, ts1[0])

	newAnimal.Height += 0.7
	err = animalService.Update(ctx, newAnimal)
	assert.NoError(t, err)

	err = animalService.Delete(ctx, id)
	assert.NoError(t, err)

	_, err = animalService.GetById(ctx, id)
	assert.Error(t, err)

	ts2, err := animalService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)
}
