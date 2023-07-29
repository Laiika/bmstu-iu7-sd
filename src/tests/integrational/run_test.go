package integrational

import (
	"context"
	"os"
	"sd/internal/config"
	"sd/internal/db/postgres"
	"sd/internal/domain/services"
	"sd/pkg/client/postgresql"
	"sd/pkg/logger"
	"testing"
)

var (
	client          postgresql.Client
	animalService   *services.AnimalService
	shelterService  *services.ShelterService
	curatorService  *services.CuratorService
	purchaseService *services.PurchaseService
)

func setup() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	var err error
	client, err = postgresql.NewClient(context.Background(), 3, &cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	animalRepo := postgres.NewAnimalRepo(client)
	animalService = services.NewAnimalService(animalRepo)

	shelterRepo := postgres.NewShelterRepo(client)
	shelterService = services.NewShelterService(shelterRepo)

	curatorRepo := postgres.NewCuratorRepo(client)
	curatorService = services.NewCuratorService(curatorRepo)

	purchaseRepo := postgres.NewPurchaseRepo(client)
	purchaseService = services.NewPurchaseService(purchaseRepo)
}

func shutdown() {
	client.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
