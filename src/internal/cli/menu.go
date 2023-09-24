package cli

import (
	"context"
	"fmt"
	"log"
	"sd/internal/config"
	mdb "sd/internal/db/mongodb"
	"sd/internal/domain/services"
	conn "sd/pkg/client/mongodb"
	"strings"
)

const menu = `
0. Зарегестрироваться
1. Просмотреть данные о себе
2. Добавить закупку
3. Добавить животное
4. Посмотреть сведения о животном
5. Посмотреть своих животных
6. Посмотреть всех животных
7. Посмотреть все приюты
8. Посмотреть сведения о приюте
9. Посмотреть сведения о закупке
10. Добавить заболение
11. Посмотреть заболевания животного
12. Выйти

Введите нужный номер: `

func start(
	animalService *services.AnimalService,
	shelterService *services.ShelterService,
	curatorService *services.CuratorService,
	purchaseService *services.PurchaseService,
	diseaseService *services.DiseaseService,
) error {
	curId := 0
	for {
		fmt.Print(menu)

		var c int
		_, err := fmt.Scan(&c)
		if err != nil {
			fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
			continue
		}

		switch c {
		case 0:
			newCurator, err := inputSignUp()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			curId, err = curatorService.Create(context.Background(), newCurator)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println("Выполнена регистрация")

		case 1:
			curator, err := curatorService.GetById(context.Background(), curId)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Имя: %s \nФамилия: %s \n Номер телефона: %s",
				curator.Name, curator.Surname, curator.PhoneNumber))

		case 2:
			newPurchase, err := inputAddPurchase()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			id, err := purchaseService.Create(context.Background(), newPurchase)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Закупка добавлена, id: %d", id))

		case 3:
			newAnimal, err := inputAddAnimal()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			id, err := animalService.Create(context.Background(), newAnimal)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Животное добавлено, id: %d", id))

		case 4:
			id, err := inputGetId()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			animal, err := animalService.GetById(context.Background(), id)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Имя: %s \nВозраст: %d \nРост: %0.3f \nВес: %0.3f \nid приюта: %d \nтип: %s \nпол: %s",
				animal.Name, animal.Age, animal.Height, animal.Weight, animal.ShelterId, animal.Type, animal.Gender))

		case 5:
			animals, err := animalService.GetCrtrAll(context.Background(), curId)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			var reply strings.Builder
			reply.WriteString(fmt.Sprintf("Найдено %d\n\n", len(animals)))
			for i := range animals {
				reply.WriteString(fmt.Sprintf("id: %d\nИмя: %s \nВозраст: %d \nРост: %0.3f \nВес: %0.3f \nid приюта: %d \nтип: %s \nпол: %s\n\n",
					animals[i].Id, animals[i].Name, animals[i].Age, animals[i].Height, animals[i].Weight, animals[i].ShelterId, animals[i].Type, animals[i].Gender))
			}
			fmt.Println(reply.String())

		case 6:
			animals, err := animalService.GetAll(context.Background())
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			var reply strings.Builder
			reply.WriteString(fmt.Sprintf("Найдено %d\n\n", len(animals)))
			for i := range animals {
				reply.WriteString(fmt.Sprintf("id: %d\nИмя: %s \nВозраст: %d \nРост: %0.3f \nВес: %0.3f \nid приюта: %d \nтип: %s \nпол: %s\n\n",
					animals[i].Id, animals[i].Name, animals[i].Age, animals[i].Height, animals[i].Weight, animals[i].ShelterId, animals[i].Type, animals[i].Gender))
			}
			fmt.Println(reply.String())

		case 7:
			shelters, err := shelterService.GetAll(context.Background())
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			var reply strings.Builder
			reply.WriteString(fmt.Sprintf("Найдено %d\n\n", len(shelters)))
			for i := range shelters {
				reply.WriteString(fmt.Sprintf("id: %d\nУлица: %s \nДом: %d\n\n",
					shelters[i].Id, shelters[i].Street, shelters[i].House))
			}
			fmt.Println(reply.String())

		case 8:
			id, err := inputGetId()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			shelter, err := shelterService.GetById(context.Background(), id)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Улица: %s \nДом: %d",
				shelter.Street, shelter.House))

		case 9:
			id, err := inputGetId()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			purchase, err := purchaseService.GetById(context.Background(), id)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Название: %s \nЧастота закупки: %s \nСтоимость: %0.3f \nДата последней: %s \nid животного: %d",
				purchase.Name, purchase.Frequency, purchase.Cost, purchase.LastDate.Format("2006-01-02"), purchase.AnimalId))

		case 10:
			newDisease, err := inputAddDisease()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			id, err := diseaseService.Create(context.Background(), newDisease)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("Заболевание добавлено, id: %d", id))
		case 11:
			id, err := inputGetId()
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			diseases, err := diseaseService.GetAnimalAll(context.Background(), id)
			if err != nil {
				fmt.Println(fmt.Sprintf("Ошибка: %s", err.Error()))
				continue
			}

			var reply strings.Builder
			reply.WriteString(fmt.Sprintf("Найдено %d\n\n", len(diseases)))
			for i := range diseases {
				isChronic := "нет"
				if diseases[i].IsChronic {
					isChronic = "да"
				}
				reply.WriteString(fmt.Sprintf("id: %d\nДиагноз: %s \nСимптомы: %s \nПричина: %s \nХроническое: %s \nid животного: %d\n\n",
					diseases[i].Id, diseases[i].Diagnosis, diseases[i].Symptoms, diseases[i].Cause, isChronic, diseases[i].AnimalId))
			}
			fmt.Println(reply.String())

		case 12:
			return nil
		}
	}
}

func Run(cfg *config.Config) {
	/*client, err := postgresql.NewClient(context.Background(), 3, &cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	animalRepo := postgres.NewAnimalRepo(client)
	animalService := services.NewAnimalService(animalRepo)

	shelterRepo := postgres.NewShelterRepo(client)
	shelterService := services.NewShelterService(shelterRepo)

	curatorRepo := postgres.NewCuratorRepo(client)
	curatorService := services.NewCuratorService(curatorRepo)

	purchaseRepo := postgres.NewPurchaseRepo(client)
	purchaseService := services.NewPurchaseService(purchaseRepo)

	diseaseRepo := postgres.NewDiseaseRepo(client)
	diseaseService := services.NewDiseaseService(diseaseRepo)*/
	db, err := conn.NewClient(context.Background(), &cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}

	animalRepo := mdb.NewAnimalRepo(db)
	animalService := services.NewAnimalService(animalRepo)

	shelterRepo := mdb.NewShelterRepo(db)
	shelterService := services.NewShelterService(shelterRepo)

	curatorRepo := mdb.NewCuratorRepo(db)
	curatorService := services.NewCuratorService(curatorRepo)

	purchaseRepo := mdb.NewPurchaseRepo(db)
	purchaseService := services.NewPurchaseService(purchaseRepo)

	diseaseRepo := mdb.NewDiseaseRepo(db)
	diseaseService := services.NewDiseaseService(diseaseRepo)

	start(animalService, shelterService, curatorService, purchaseService, diseaseService)
}
