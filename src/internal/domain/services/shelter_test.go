package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"sd/internal/domain/services/mocks"
	"testing"
)

func TestShelterService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIShelterRepo(ctl)
	service := NewShelterService(repo)

	id := 11
	expectedShelter := &entities.Shelter{
		Id:     id,
		Street: "Первомайская",
		House:  10,
	}

	repo.EXPECT().GetById(gomock.Any(), id).Return(expectedShelter, nil)
	shelter, err := service.GetById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedShelter, shelter)
}

func TestShelterService_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIShelterRepo(ctl)
	service := NewShelterService(repo)

	expectedShelters := entities.Shelters{
		&entities.Shelter{
			Id:     1,
			Street: "Первомайская",
			House:  10,
		},
		&entities.Shelter{
			Id:     2,
			Street: "Анникова",
			House:  11,
		},
	}

	repo.EXPECT().GetAll(gomock.Any()).Return(expectedShelters, nil)
	shelters, err := service.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedShelters, shelters)
}

func TestShelterService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIShelterRepo(ctl)
	service := NewShelterService(repo)

	shelter := &entities.Shelter{
		Street: "Первомайская",
		House:  10,
	}

	id := 1
	repo.EXPECT().Create(gomock.Any(), shelter).Return(id, nil)
	id2, err := service.Create(context.Background(), shelter)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}

func TestShelterService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIShelterRepo(ctl)
	service := NewShelterService(repo)

	shelter := &entities.Shelter{
		Id:     1,
		Street: "Первомайская",
		House:  10,
	}

	repo.EXPECT().Update(gomock.Any(), shelter).Return(nil)
	err := service.Update(context.Background(), shelter)
	assert.NoError(t, err)
}

func TestShelterService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIShelterRepo(ctl)
	service := NewShelterService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
