package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"sd/internal/domain/services/mocks"
	"testing"
	"time"
)

func TestPurchaseService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	id := 11
	expectedPurchase := &entities.Purchase{
		Id:        id,
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	repo.EXPECT().GetById(gomock.Any(), id).Return(expectedPurchase, nil)
	purchase, err := service.GetById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedPurchase, purchase)
}

func TestPurchaseService_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	expectedPurchases := entities.Purchases{
		&entities.Purchase{
			Id:        1,
			Name:      "Royal Canin Sensible 33",
			Frequency: "раз в 2 недели",
			Cost:      403,
			LastDate:  time.Now(),
			AnimalId:  14,
		},
		&entities.Purchase{
			Id:        2,
			Name:      "Royal Canin Sterilised в желе",
			Frequency: "раз в 2 недели",
			Cost:      800,
			LastDate:  time.Now(),
			AnimalId:  14,
		},
	}

	repo.EXPECT().GetAll(gomock.Any()).Return(expectedPurchases, nil)
	purchases, err := service.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedPurchases, purchases)
}

func TestPurchaseService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	purchase := &entities.Purchase{
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	id := 1
	repo.EXPECT().Create(gomock.Any(), purchase).Return(id, nil)
	id2, err := service.Create(context.Background(), purchase)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}

func TestPurchaseService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	purchase := &entities.Purchase{
		Id:        1,
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	repo.EXPECT().Update(gomock.Any(), purchase).Return(nil)
	err := service.Update(context.Background(), purchase)
	assert.NoError(t, err)
}

func TestPurchaseService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
