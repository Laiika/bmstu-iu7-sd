package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/dto"
	"sd/internal/domain/entities"
	"sd/internal/domain/services/mocks"
	"testing"
)

func TestAnimalService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	id := 11
	expectedAnimal := &entities.Animal{
		Id:        id,
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}

	repo.EXPECT().GetById(gomock.Any(), id).Return(expectedAnimal, nil)
	animal, err := service.GetById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnimal, animal)
}

func TestAnimalService_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	expectedAnimals := entities.Animals{
		&entities.Animal{
			Id:        1,
			Name:      "Бобик",
			Age:       10,
			Height:    3.4,
			Weight:    5.5,
			ShelterId: 10,
			Type:      "собака",
			Gender:    "мужской",
		},
		&entities.Animal{
			Id:        2,
			Name:      "Шарик",
			Age:       6,
			Height:    4.4,
			Weight:    6.5,
			ShelterId: 10,
			Type:      "собака",
			Gender:    "мужской",
		},
	}

	repo.EXPECT().GetAll(gomock.Any()).Return(expectedAnimals, nil)
	animals, err := service.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedAnimals, animals)
}

func TestAnimalService_GetCrtrAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	expectedAnimals := entities.Animals{
		&entities.Animal{
			Id:        1,
			Name:      "Бобик",
			Age:       10,
			Height:    3.4,
			Weight:    5.5,
			ShelterId: 10,
			Type:      "собака",
			Gender:    "мужской",
		},
		&entities.Animal{
			Id:        2,
			Name:      "Шарик",
			Age:       6,
			Height:    4.4,
			Weight:    6.5,
			ShelterId: 10,
			Type:      "собака",
			Gender:    "мужской",
		},
	}
	crtrId := 15

	repo.EXPECT().GetCrtrAll(gomock.Any(), crtrId).Return(expectedAnimals, nil)
	animals, err := service.GetCrtrAll(context.Background(), crtrId)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnimals, animals)
}

func TestAnimalService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	dto := &dto.CreateAnimal{
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}

	repo.EXPECT().Create(gomock.Any(), dto).Return(nil)
	err := service.Create(context.Background(), dto)
	assert.NoError(t, err)
}

func TestAnimalService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	dto := &dto.UpdateAnimal{
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}
	id := 1

	repo.EXPECT().Update(gomock.Any(), id, dto).Return(nil)
	err := service.Update(context.Background(), id, dto)
	assert.NoError(t, err)
}

func TestAnimalService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
