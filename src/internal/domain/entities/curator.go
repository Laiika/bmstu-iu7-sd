package entities

import (
	"regexp"
	"sd/internal/apperrors"
)

type Curator struct {
	Id          int
	ChatId      string
	Name        string
	Surname     string
	PhoneNumber string
}

type Curators []*Curator

func (c *Curator) isValidNumber() error {
	match, _ := regexp.MatchString(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`, c.PhoneNumber)

	if !match {
		return apperrors.ErrInvalidPhoneNumber
	}
	return nil
}

func (c *Curator) IsValid() error {
	err := c.isValidNumber()

	switch {
	case err != nil:
		break
	case c.Name == "":
		err = apperrors.ErrInvalidName
	case c.Surname == "":
		err = apperrors.ErrInvalidSurname
	}

	return err
}
