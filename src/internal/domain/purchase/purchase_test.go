package purchase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPurchaseService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	id := 11
	expectedPurchase := &Purchase{
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
	repo := NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	expectedPurchases := Purchases{
		&Purchase{
			Id:        1,
			Name:      "Royal Canin Sensible 33",
			Frequency: "раз в 2 недели",
			Cost:      403,
			LastDate:  time.Now(),
			AnimalId:  14,
		},
		&Purchase{
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
	repo := NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	dto := &CreatePurchase{
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	expectedPurchase := &Purchase{
		Id:        1,
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	repo.EXPECT().Create(gomock.Any(), dto).Return(expectedPurchase, nil)
	purchase, err := service.Create(context.Background(), dto)
	assert.NoError(t, err)
	assert.Equal(t, expectedPurchase, purchase)
}

func TestPurchaseService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	dto := &UpdatePurchase{
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	id := 1
	expectedPurchase := &Purchase{
		Id:        id,
		Name:      "Royal Canin Sensible 33",
		Frequency: "раз в 2 недели",
		Cost:      403,
		LastDate:  time.Now(),
		AnimalId:  14,
	}

	repo.EXPECT().Update(gomock.Any(), id, dto).Return(expectedPurchase, nil)
	purchase, err := service.Update(context.Background(), id, dto)
	assert.NoError(t, err)
	assert.Equal(t, expectedPurchase, purchase)
}

func TestPurchaseService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIPurchaseRepo(ctl)
	service := NewPurchaseService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
