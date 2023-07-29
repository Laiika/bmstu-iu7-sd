package entities

import (
	"sd/internal/apperrors"
	"time"
)

type Purchase struct {
	Id        int
	Name      string
	Frequency string
	Cost      float64
	LastDate  time.Time
	AnimalId  int
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
