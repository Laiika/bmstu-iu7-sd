package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"testing"
)

func TestCuratorService(t *testing.T) {
	ctx := context.Background()
	err := TruncateTables(client, ctx)
	if err != nil {
		return
	}

	newCurator := &entities.Curator{
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}
	id, err := curatorService.Create(ctx, newCurator)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	newCurator.Id = id
	t1, err := curatorService.GetById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, newCurator, t1)

	ts1, err := curatorService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, newCurator, ts1[0])

	newCurator.PhoneNumber = "+79045151234"
	err = curatorService.Update(ctx, newCurator)
	assert.NoError(t, err)

	err = curatorService.Delete(ctx, id)
	assert.NoError(t, err)

	_, err = curatorService.GetById(ctx, id)
	assert.Error(t, err)

	ts2, err := curatorService.GetAll(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)
}
