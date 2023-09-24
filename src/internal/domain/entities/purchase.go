package entities

import (
	"sd/internal/apperrors"
	"time"
)

type Purchase struct {
	Id        int       `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Frequency string    `bson:"frequency"`
	Cost      float64   `bson:"cost"`
	LastDate  time.Time `bson:"last_date"`
	AnimalId  int       `bson:"animal_id"`
}

type Purchases []*Purchase

func (p *Purchase) IsValid() error {
	var err error

	switch {
	case p.Name == "":
		err = apperrors.ErrInvalidName
	case p.AnimalId <= 0:
		err = apperrors.ErrInvalidReferenceId
	}

	return err
}
