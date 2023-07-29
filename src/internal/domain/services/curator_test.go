package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sd/internal/domain/entities"
	"sd/internal/domain/services/mocks"
	"testing"
)

func TestCuratorService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	id := 11
	expectedCurator := &entities.Curator{
		Id:          id,
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	repo.EXPECT().GetById(gomock.Any(), id).Return(expectedCurator, nil)
	curator, err := service.GetById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedCurator, curator)
}

func TestCuratorService_GetByChatId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	chatId := "85085228"
	expectedCurator := &entities.Curator{
		Id:          11,
		ChatId:      chatId,
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	repo.EXPECT().GetByChatId(gomock.Any(), chatId).Return(expectedCurator, nil)
	curator, err := service.GetByChatId(context.Background(), chatId)
	assert.NoError(t, err)
	assert.Equal(t, expectedCurator, curator)
}

func TestCuratorService_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	expectedCurators := entities.Curators{
		&entities.Curator{
			Id:          1,
			ChatId:      "85085228",
			Name:        "Арина",
			Surname:     "Иванова",
			PhoneNumber: "+79891143454",
		},
		&entities.Curator{
			Id:          2,
			ChatId:      "85033128",
			Name:        "Анастасия",
			Surname:     "Орехова",
			PhoneNumber: "+79891003417",
		},
	}

	repo.EXPECT().GetAll(gomock.Any()).Return(expectedCurators, nil)
	curators, err := service.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedCurators, curators)
}

func TestCuratorService_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	curator := &entities.Curator{
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	id := 1
	repo.EXPECT().Create(gomock.Any(), curator).Return(id, nil)
	id2, err := service.Create(context.Background(), curator)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}

func TestCuratorService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	curator := &entities.Curator{
		Id:          1,
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	repo.EXPECT().Update(gomock.Any(), curator).Return(nil)
	err := service.Update(context.Background(), curator)
	assert.NoError(t, err)
}

func TestCuratorService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mocks.NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
