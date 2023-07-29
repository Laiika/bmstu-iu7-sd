package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"testing"
)

func TestShelterService(t *testing.T) {
	ctx := context.Background()
	err := TruncateTables(client, ctx)
	if err != nil {
		return
	}

	newShelter := &entities.Shelter{
		Street: "Первомайская",
		House:  10,
	}
	id, err := shelterService.Create(ctx, newShelter)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	newShelter.Id = id
	t1, err := shelterService.GetById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, newShelter, t1)

	ts1, err := shelterService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, newShelter, ts1[0])

	newShelter.House = 24
	err = shelterService.Update(ctx, newShelter)
	assert.NoError(t, err)

	err = shelterService.Delete(ctx, id)
	assert.NoError(t, err)

	_, err = shelterService.GetById(ctx, id)
	assert.Error(t, err)

	ts2, err := shelterService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)
}
