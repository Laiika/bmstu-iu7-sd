package cli

import (
	"fmt"
	"sd/internal/domain/entities"
	"time"
)

func inputSignUp() (*entities.Curator, error) {
	var err error
	newCurator := &entities.Curator{
		ChatId: "123",
	}

	var name string
	fmt.Println("Введите имя:")
	if _, err = fmt.Scan(&name); err != nil {
		return nil, err
	}
	newCurator.Name = name

	var surname string
	fmt.Println("Введите фамилию:")
	if _, err = fmt.Scan(&surname); err != nil {
		return nil, err
	}
	newCurator.Surname = surname

	var phone string
	fmt.Println("Введите номер телефона:")
	if _, err = fmt.Scan(&phone); err != nil {
		return nil, err
	}
	newCurator.PhoneNumber = phone

	return newCurator, nil
}

func inputAddPurchase() (*entities.Purchase, error) {
	var err error
	newPurchase := &entities.Purchase{}

	var name string
	fmt.Println("Введите название:")
	if _, err = fmt.Scan(&name); err != nil {
		return nil, err
	}
	newPurchase.Name = name

	var fr string
	fmt.Println("Введите частоту закупки:")
	if _, err = fmt.Scan(&fr); err != nil {
		return nil, err
	}
	newPurchase.Frequency = fr

	var cost float64
	fmt.Println("Введите стоимость:")
	if _, err = fmt.Scan(&cost); err != nil {
		return nil, err
	}
	newPurchase.Cost = cost

	var date string
	fmt.Println("Введите дату последней закупки (ГГГГ-ММ-ДД):")
	if _, err = fmt.Scan(&date); err != nil {
		return nil, err
	}
	newPurchase.LastDate, err = time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	var an int
	fmt.Println("Введите id животного:")
	if _, err = fmt.Scan(&an); err != nil {
		return nil, err
	}
	newPurchase.AnimalId = an

	return newPurchase, nil
}

func inputAddAnimal() (*entities.Animal, error) {
	var err error
	newAnimal := &entities.Animal{}

	var name string
	fmt.Println("Введите кличку:")
	if _, err = fmt.Scan(&name); err != nil {
		return nil, err
	}
	newAnimal.Name = name

	var age int
	fmt.Println("Введите возраст:")
	if _, err = fmt.Scan(&age); err != nil {
		return nil, err
	}
	newAnimal.Age = age

	var h float64
	fmt.Println("Введите рост:")
	if _, err = fmt.Scan(&h); err != nil {
		return nil, err
	}
	newAnimal.Height = h

	var w float64
	fmt.Println("Введите вес:")
	if _, err = fmt.Scan(&w); err != nil {
		return nil, err
	}
	newAnimal.Weight = w

	var sh int
	fmt.Println("Введите id приюта:")
	if _, err = fmt.Scan(&sh); err != nil {
		return nil, err
	}
	newAnimal.ShelterId = sh

	var t string
	fmt.Println("Введите тип (кошка/собака):")
	if _, err = fmt.Scan(&t); err != nil {
		return nil, err
	}
	newAnimal.Type = t

	var g string
	fmt.Println("Введите пол (мужской/женский):")
	if _, err = fmt.Scan(&g); err != nil {
		return nil, err
	}
	newAnimal.Gender = g

	return newAnimal, nil
}

func inputAddDisease() (*entities.Disease, error) {
	var err error
	newDisease := &entities.Disease{}

	var diagnosis string
	fmt.Println("Введите диагноз:")
	if _, err = fmt.Scan(&diagnosis); err != nil {
		return nil, err
	}
	newDisease.Diagnosis = diagnosis

	var symptoms string
	fmt.Println("Введите симптомы:")
	if _, err = fmt.Scan(&symptoms); err != nil {
		return nil, err
	}
	newDisease.Symptoms = symptoms

	var cause string
	fmt.Println("Введите причину заболевания:")
	if _, err = fmt.Scan(&cause); err != nil {
		return nil, err
	}
	newDisease.Cause = cause

	var isChronic string
	fmt.Println("Введите, является ли заболевание хроническим (да/нет):")
	if _, err = fmt.Scan(&isChronic); err != nil {
		return nil, err
	}

	newDisease.IsChronic = false
	if isChronic == "да" {
		newDisease.IsChronic = true
	}

	var id int
	fmt.Println("Введите id животного:")
	if _, err = fmt.Scan(&id); err != nil {
		return nil, err
	}
	newDisease.AnimalId = id

	return newDisease, nil
}

func inputGetId() (int, error) {
	var err error

	var id int
	fmt.Println("Введите id:")
	if _, err = fmt.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
