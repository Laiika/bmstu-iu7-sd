package entities

type Shelter struct {
	Id     int
	Street string
	House  int
}

type Shelters []*Shelter
