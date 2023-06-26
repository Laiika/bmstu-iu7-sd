package animal

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnimalService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	id := 11
	expectedAnimal := &Animal{
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
	repo := NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	expectedAnimals := Animals{
		&Animal{
			Id:        1,
			Name:      "Бобик",
			Age:       10,
			Height:    3.4,
			Weight:    5.5,
			ShelterId: 10,
			Type:      "собака",
			Gender:    "мужской",
		},
		&Animal{
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
	repo := NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	expectedAnimals := Animals{
		&Animal{
			Id:        1,
			Name:      "Бобик",
			Age:       10,
			Height:    3.4,
			Weight:    5.5,
			ShelterId: 10,
			Type:      "собака",
			Gender:    "мужской",
		},
		&Animal{
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
	repo := NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	dto := &CreateAnimal{
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}

	expectedAnimal := &Animal{
		Id:        1,
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}

	repo.EXPECT().Create(gomock.Any(), dto).Return(expectedAnimal, nil)
	animal, err := service.Create(context.Background(), dto)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnimal, animal)
}

func TestAnimalService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	dto := &UpdateAnimal{
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}

	id := 1
	expectedAnimal := &Animal{
		Id:        id,
		Name:      "Бобик",
		Age:       10,
		Height:    3.4,
		Weight:    5.5,
		ShelterId: 10,
		Type:      "собака",
		Gender:    "мужской",
	}

	repo.EXPECT().Update(gomock.Any(), id, dto).Return(expectedAnimal, nil)
	animal, err := service.Update(context.Background(), id, dto)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnimal, animal)
}

func TestAnimalService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIAnimalRepo(ctl)
	service := NewAnimalService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
