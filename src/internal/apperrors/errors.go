package apperrors

import "errors"

var (
	ErrInvalidReferenceId = errors.New("невалидный id")

	ErrInvalidName   = errors.New("невалидное имя")
	ErrInvalidAge    = errors.New("невалидный возраст")
	ErrInvalidType   = errors.New("невалидный тип")
	ErrInvalidGender = errors.New("невалидный пол")

	ErrInvalidSurname     = errors.New("невалидная фамилия")
	ErrInvalidPhoneNumber = errors.New("невалидный номер телефона")

	ErrInvalidStreet      = errors.New("невалидная улица")
	ErrInvalidHouseNumber = errors.New("невалидный номер дома")

	ErrEntityExists   = errors.New("такая сущность уже существует")
	ErrInternal       = errors.New("внутренняя ошибка базы данных")
	ErrEntityNotFound = errors.New("сущность не найдена")
)
