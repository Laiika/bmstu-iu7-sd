package entities

import "sd/internal/apperrors"

type Shelter struct {
	Id     int    `bson:"_id,omitempty"`
	Street string `bson:"street"`
	House  int    `bson:"house"`
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
