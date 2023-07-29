package entities

import "sd/internal/apperrors"

type Shelter struct {
	Id     int
	Street string
	House  int
}

type Shelters []*Shelter

func (sh *Shelter) IsValid() error {
	var err error

	switch {
	case sh.Street == "":
		err = apperrors.ErrInvalidStreet
	case sh.House <= 0:
		err = apperrors.ErrInvalidHouseNumber
	}

	return err
}
