package dto

import "time"

type CreatePurchase struct {
	Name      string
	Frequency string
	Cost      float64
	LastDate  time.Time
	AnimalId  int
}

type UpdatePurchase struct {
	Name      string
	Frequency string
	Cost      float64
	LastDate  time.Time
	AnimalId  int
}
