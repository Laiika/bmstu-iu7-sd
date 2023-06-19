package curator

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCuratorService_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	id := 11
	expectedCurator := &Curator{
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

func TestCuratorService_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	expectedCurators := Curators{
		&Curator{
			Id:          1,
			ChatId:      "85085228",
			Name:        "Арина",
			Surname:     "Иванова",
			PhoneNumber: "+79891143454",
		},
		&Curator{
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
	repo := NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	dto := &CreateCurator{
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	expectedCurator := &Curator{
		Id:          1,
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	repo.EXPECT().Create(gomock.Any(), dto).Return(expectedCurator, nil)
	curator, err := service.Create(context.Background(), dto)
	assert.NoError(t, err)
	assert.Equal(t, expectedCurator, curator)
}

func TestCuratorService_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	dto := &UpdateCurator{
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	id := 1
	expectedCurator := &Curator{
		Id:          id,
		ChatId:      "85085228",
		Name:        "Арина",
		Surname:     "Иванова",
		PhoneNumber: "+79891143454",
	}

	repo.EXPECT().Update(gomock.Any(), id, dto).Return(expectedCurator, nil)
	curator, err := service.Update(context.Background(), id, dto)
	assert.NoError(t, err)
	assert.Equal(t, expectedCurator, curator)
}

func TestCuratorService_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockICuratorRepo(ctl)
	service := NewCuratorService(repo)

	id := 1
	repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	err := service.Delete(context.Background(), id)
	assert.NoError(t, err)
}
