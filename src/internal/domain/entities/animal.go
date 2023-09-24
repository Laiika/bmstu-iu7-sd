package entities

import (
	"sd/internal/apperrors"
	"strings"
)

type Animal struct {
	Id        int     `bson:"_id,omitempty"`
	Name      string  `bson:"name"`
	Age       int     `bson:"age"`
	Height    float64 `bson:"height"`
	Weight    float64 `bson:"weight"`
	ShelterId int     `bson:"shelter_id"`
	Type      string  `bson:"type"`
	Gender    string  `bson:"gender"`
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
