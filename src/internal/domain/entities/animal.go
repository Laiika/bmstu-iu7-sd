package entities

import (
	"sd/internal/apperrors"
	"strings"
)

type Animal struct {
	Id        int
	Name      string
	Age       int
	Height    float64
	Weight    float64
	ShelterId int
	Type      string
	Gender    string
}

type Animals []*Animal

func (an *Animal) IsValid() error {
	var err error

	switch {
	case an.Name == "":
		err = apperrors.ErrInvalidName
	case an.Age <= 0:
		err = apperrors.ErrInvalidAge
	case an.ShelterId <= 0:
		err = apperrors.ErrInvalidReferenceId
	case !strings.EqualFold(an.Type, "собака") && !strings.EqualFold(an.Type, "кошка"):
		err = apperrors.ErrInvalidType
	case !strings.EqualFold(an.Gender, "мужской") && !strings.EqualFold(an.Gender, "женский"):
		err = apperrors.ErrInvalidGender
	}

	return err
}
