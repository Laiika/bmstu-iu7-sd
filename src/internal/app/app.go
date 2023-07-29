package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sd/internal/config"
	"sd/internal/controllers"
	"sd/internal/db/postgres"
	"sd/internal/domain/entities"
	"sd/internal/domain/services"
	"sd/pkg/client/postgresql"
	"sd/pkg/logger"
	"strconv"
	"time"
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Зарегестрироваться"),
		tgbotapi.NewKeyboardButton("Просмотреть данные о себе"),
		tgbotapi.NewKeyboardButton("Добавить закупку"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Добавить животное"),
		tgbotapi.NewKeyboardButton("Добавить животное к себе"),
		tgbotapi.NewKeyboardButton("Перестать быть куратором животного"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Посмотреть сведения о животном"),
		tgbotapi.NewKeyboardButton("Посмотреть своих животных"),
		tgbotapi.NewKeyboardButton("Посмотреть всех животных"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Посмотреть все приюты"),
		tgbotapi.NewKeyboardButton("Посмотреть сведения о приюте"),
		tgbotapi.NewKeyboardButton("Посмотреть сведения о закупке"),
	),
)

func singUp(chatId int64, controller *controllers.Controller, data []string) string {
	newCurator := &entities.Curator{
		ChatId:      strconv.FormatInt(chatId, 10),
		Name:        data[0],
		Surname:     data[1],
		PhoneNumber: data[2],
	}
	_, err := controller.AddCurator(context.Background(), newCurator)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return "Выполнена регистрация"
}

func addAnimal(controller *controllers.Controller, data []string) string {
	age, _ := strconv.Atoi(data[1])
	height, _ := strconv.ParseFloat(data[2], 64)
	weight, _ := strconv.ParseFloat(data[3], 64)
	shId, _ := strconv.Atoi(data[4])

	newAnimal := &entities.Animal{
		Name:      data[0],
		Age:       age,
		Height:    height,
		Weight:    weight,
		ShelterId: shId,
		Type:      data[5],
		Gender:    data[6],
	}
	id, err := controller.AddAnimal(context.Background(), newAnimal)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return fmt.Sprintf("Животное добавлено, id: %d", id)
}

func addPurchase(controller *controllers.Controller, data []string) string {
	cost, _ := strconv.ParseFloat(data[2], 64)
	lastDate, _ := time.Parse("2006-01-02", data[3])
	anId, _ := strconv.Atoi(data[4])

	newPurchase := &entities.Purchase{
		Name:      data[0],
		Frequency: data[1],
		Cost:      cost,
		LastDate:  lastDate,
		AnimalId:  anId,
	}
	id, err := controller.AddPurchase(context.Background(), newPurchase)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return fmt.Sprintf("Закупка добавлена, id: %d", id)
}

func getInfoAboutPurchase(controller *controllers.Controller, data []string) string {
	id, _ := strconv.Atoi(data[0])

	purchase, err := controller.GetPurchase(context.Background(), id)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return fmt.Sprintf("Название: %s \nЧастота закупки: %s \nСтоимость: %0.3f \nДата последней: %s \nid животного: %d",
		purchase.Name, purchase.Frequency, purchase.Cost, purchase.LastDate.Format("2006-01-02"), purchase.AnimalId)
}

func getInfoAboutAnimal(controller *controllers.Controller, data []string) string {
	id, _ := strconv.Atoi(data[0])

	animal, err := controller.GetAnimal(context.Background(), id)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return fmt.Sprintf("Имя: %s \nВозраст: %d \nРост: %0.3f \nВес: %0.3f \nid приюта: %d \nтип: %s \nпол: %s",
		animal.Name, animal.Age, animal.Height, animal.Weight, animal.ShelterId, animal.Type, animal.Gender)
}

func getInfoAboutShelter(controller *controllers.Controller, data []string) string {
	id, _ := strconv.Atoi(data[0])

	shelter, err := controller.GetShelter(context.Background(), id)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return fmt.Sprintf("Улица: %s \nДом: %d",
		shelter.Street, shelter.House)
}

func getInfoAboutCurator(chatId int64, controller *controllers.Controller) string {
	curator, err := controller.GetCuratorByChatId(context.Background(), strconv.FormatInt(chatId, 10))
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	return fmt.Sprintf("Имя: %s \nФамилия: %s \n Номер телефона: %s",
		curator.Name, curator.Surname, curator.PhoneNumber)
}

func printCrtrAnimals(chatId int64, controller *controllers.Controller) string {
	curator, err := controller.GetCuratorByChatId(context.Background(), strconv.FormatInt(chatId, 10))
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	animals, err := controller.GetAllCrtrAnimals(context.Background(), curator.Id)
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	reply := fmt.Sprintf("Найдено %d\n\n", len(animals))
	for i := range animals {
		reply += fmt.Sprintf("id: %d\nИмя: %s \nВозраст: %d \nРост: %0.3f \nВес: %0.3f \nid приюта: %d \nтип: %s \nпол: %s\n\n",
			animals[i].Id, animals[i].Name, animals[i].Age, animals[i].Height, animals[i].Weight, animals[i].ShelterId, animals[i].Type, animals[i].Gender)
	}
	return reply
}

func printAllAnimals(controller *controllers.Controller) string {
	animals, err := controller.GetAllAnimals(context.Background())
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	reply := fmt.Sprintf("Найдено %d\n\n", len(animals))
	for i := range animals {
		reply += fmt.Sprintf("id: %d\nИмя: %s \nВозраст: %d \nРост: %0.3f \nВес: %0.3f \nid приюта: %d \nтип: %s \nпол: %s\n\n",
			animals[i].Id, animals[i].Name, animals[i].Age, animals[i].Height, animals[i].Weight, animals[i].ShelterId, animals[i].Type, animals[i].Gender)
	}
	return reply
}

func printAllShelters(controller *controllers.Controller) string {
	shelters, err := controller.GetAllShelters(context.Background())
	if err != nil {
		return fmt.Sprintf("Ошибка: %s", err.Error())
	}

	reply := fmt.Sprintf("Найдено %d\n\n", len(shelters))
	for i := range shelters {
		reply += fmt.Sprintf("id: %d\nУлица: %s \nДом: %d\n\n",
			shelters[i].Id, shelters[i].Street, shelters[i].House)
	}
	return reply
}

func runCommand(chatId int64, command string, controller *controllers.Controller, data []string) string {
	result := ""

	switch command {
	case "Зарегестрироваться":
		result = singUp(chatId, controller, data)
	case "Добавить животное":
		result = addAnimal(controller, data)
	case "Добавить закупку":
		result = addPurchase(controller, data)
	case "Посмотреть сведения о закупке":
		result = getInfoAboutPurchase(controller, data)
	case "Посмотреть сведения о животном":
		result = getInfoAboutAnimal(controller, data)
	case "Посмотреть сведения о приюте":
		result = getInfoAboutShelter(controller, data)
	}

	return result
}

func singUpMessages() []string {
	var ms = []string{
		"Введите имя",
		"Введите фамилию",
		"Введите номер телефона",
	}
	return ms
}

func getInfoAboutMessages() []string {
	var ms = []string{
		"Введите id",
	}
	return ms
}

func addAnimalMessages() []string {
	var ms = []string{
		"Введите кличку",
		"Введите возраст",
		"Введите рост",
		"Введите вес",
		"Введите id приюта",
		"Введите тип (кошка/собака)",
		"Введите пол (мужской/женский)",
	}
	return ms
}

func addPurchaseMessages() []string {
	var ms = []string{
		"Введите название",
		"Введите частоту закупки",
		"Введите стоимость",
		"Введите дату последней закупки (ГГГГ-ММ-ДД)",
		"Введите id животного",
	}
	return ms
}

func handleUpdates(controller *controllers.Controller, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, log *logger.Logger) {
	var command string
	var ms []string
	var data = make([]string, 10)
	num := 0

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		chatId := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatId, update.Message.Text)
		log.Info(msg.Text)
		msg.ReplyMarkup = keyboard

		if num > 0 {
			if num < len(ms) {
				data[num-1] = update.Message.Text
				msg.Text = ms[num]
				num++
			} else {
				data[num-1] = update.Message.Text
				msg.Text = runCommand(chatId, command, controller, data)
				num = 0
			}
		} else {
			switch update.Message.Text {
			case "Зарегестрироваться":
				command = update.Message.Text
				ms = singUpMessages()
				msg.Text = ms[num]
				num = 1
			case "Просмотреть данные о себе":
				msg.Text = getInfoAboutCurator(chatId, controller)
			case "Добавить животное":
				command = update.Message.Text
				ms = addAnimalMessages()
				msg.Text = ms[num]
				num = 1
			case "Добавить животное к себе":

			case "Перестать быть куратором животного":

			case "Посмотреть сведения о животном":
				command = update.Message.Text
				ms = getInfoAboutMessages()
				msg.Text = ms[num]
				num = 1
			case "Посмотреть своих животных":
				msg.Text = printCrtrAnimals(chatId, controller)
			case "Посмотреть всех животных":
				msg.Text = printAllAnimals(controller)
			case "Посмотреть все приюты":
				msg.Text = printAllShelters(controller)
			case "Посмотреть сведения о приюте":
				command = update.Message.Text
				ms = getInfoAboutMessages()
				msg.Text = ms[num]
				num = 1
			case "Посмотреть сведения о закупке":
				command = update.Message.Text
				ms = getInfoAboutMessages()
				msg.Text = ms[num]
				num = 1
			case "Добавить закупку":
				command = update.Message.Text
				ms = addPurchaseMessages()
				msg.Text = ms[num]
				num = 1
			default:
				msg.Text = "Команда не найдена"
				log.Warn("Неверная команда")
			}
		}

		log.Info(msg.Text)

		if _, err := bot.Send(msg); err != nil {
			log.Fatal("Ошибка запроса")
		}
	}
}

func startBot(controller *controllers.Controller, cfg *config.Config, log *logger.Logger) {
	token := cfg.TelegramToken
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = 30
	updates := bot.GetUpdatesChan(updateCfg)

	handleUpdates(controller, updates, bot, log)
}

func Run(cfg *config.Config, log *logger.Logger) {
	client, err := postgresql.NewClient(context.Background(), 3, &cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Info("удалось подключиться к бд")

	animalRepo := postgres.NewAnimalRepo(client)
	animalService := services.NewAnimalService(animalRepo)

	shelterRepo := postgres.NewShelterRepo(client)
	shelterService := services.NewShelterService(shelterRepo)

	curatorRepo := postgres.NewCuratorRepo(client)
	curatorService := services.NewCuratorService(curatorRepo)

	purchaseRepo := postgres.NewPurchaseRepo(client)
	purchaseService := services.NewPurchaseService(purchaseRepo)

	controller := controllers.NewController(curatorService, animalService, shelterService, purchaseService)
	startBot(controller, cfg, log)
}
