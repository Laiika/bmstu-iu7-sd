package purchase

import "time"

type Purchase struct {
	Id        int
	Name      string
	Frequency string
	Cost      float64
	LastDate  time.Time
	AnimalId  int
}

type Purchases []*Purchase

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
