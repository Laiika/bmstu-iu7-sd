package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"sd/internal/domain/services/mocks"
	"testing"
)

func TestDiseaseService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIDiseaseRepo(ctl)
	service := NewDiseaseService(repo)

	id := 11
	expectedDisease := &entities.Disease{
		Id:        id,
		Diagnosis: "почечная недостаточность",
		Symptoms:  "рвота, вялость",
		Cause:     "дегидратация",
		IsChronic: false,
		AnimalId:  id,
	}

	repo.EXPECT().GetById(gomock.Any(), id).Return(expectedDisease, nil)
	disease, err := service.GetById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedDisease, disease)
}

func TestDiseaseService_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIDiseaseRepo(ctl)
	service := NewDiseaseService(repo)

	expectedDiseases := entities.Diseases{
		&entities.Disease{
			Id:        1,
			Diagnosis: "почечная недостаточность",
			Symptoms:  "рвота, вялость",
			Cause:     "дегидратация",
			IsChronic: false,
			AnimalId:  10,
		},
		&entities.Disease{
			Id:        2,
			Diagnosis: "цистит",
			Symptoms:  "частое мочеиспускание",
			Cause:     "мочекаменная болезнь",
			IsChronic: true,
			AnimalId:  15,
		},
	}

	repo.EXPECT().GetAll(gomock.Any()).Return(expectedDiseases, nil)
	diseases, err := service.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedDiseases, diseases)
}

func TestDiseaseService_GetAnimalAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIDiseaseRepo(ctl)
	service := NewDiseaseService(repo)

	expectedDiseases := entities.Diseases{
		&entities.Disease{
			Id:        1,
			Diagnosis: "почечная недостаточность",
			Symptoms:  "рвота, вялость",
			Cause:     "дегидратация",
			IsChronic: false,
			AnimalId:  15,
		},
		&entities.Disease{
			Id:        2,
			Diagnosis: "цистит",
			Symptoms:  "частое мочеиспускание",
			Cause:     "мочекаменная болезнь",
			IsChronic: true,
			AnimalId:  15,
		},
	}
	anId := 15

	repo.EXPECT().GetAnimalAll(gomock.Any(), anId).Return(expectedDiseases, nil)
	diseases, err := service.GetAnimalAll(context.Background(), anId)
	assert.NoError(t, err)
	assert.Equal(t, expectedDiseases, diseases)
}

func TestDiseaseService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIDiseaseRepo(ctl)
	service := NewDiseaseService(repo)

	disease := &entities.Disease{
		Diagnosis: "почечная недостаточность",
		Symptoms:  "рвота, вялость",
		Cause:     "дегидратация",
		IsChronic: false,
		AnimalId:  15,
	}

	id := 1
	repo.EXPECT().Create(gomock.Any(), disease).Return(id, nil)
	id2, err := service.Create(context.Background(), disease)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}

func TestDiseaseService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIDiseaseRepo(ctl)
	service := NewDiseaseService(repo)

	disease := &entities.Disease{
		Id:        1,
		Diagnosis: "почечная недостаточность",
		Symptoms:  "рвота, вялость",
		Cause:     "дегидратация",
		IsChronic: false,
		AnimalId:  15,
	}

	repo.EXPECT().Update(gomock.Any(), disease).Return(nil)
	err := service.Update(context.Background(), disease)
	assert.NoError(t, err)
}

func TestDiseaseService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIDiseaseRepo(ctl)
	service := NewDiseaseService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
