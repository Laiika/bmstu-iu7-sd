package entities

import (
	"sd/internal/apperrors"
)

type Disease struct {
	Id        int    `bson:"_id,omitempty"`
	Diagnosis string `bson:"diagnosis"`
	Symptoms  string `bson:"symptoms"`
	Cause     string `bson:"cause"`
	IsChronic bool   `bson:"is_chronic"`
	AnimalId  int    `bson:"animal_id"`
}

type Diseases []*Disease

func (d *Disease) IsValid() error {
	var err error

	switch {
	case d.Diagnosis == "":
		err = apperrors.ErrInvalidDiagnosis
	case d.Symptoms == "":
		err = apperrors.ErrInvalidSymptoms
	case d.AnimalId <= 0:
		err = apperrors.ErrInvalidReferenceId
	}

	return err
}
