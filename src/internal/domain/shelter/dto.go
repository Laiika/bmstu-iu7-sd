package shelter

type Shelter struct {
	Id     int
	Street string
	House  int
}

type Shelters []*Shelter

type CreateShelter struct {
	Street string
	House  int
}

type UpdateShelter struct {
	Street string
	House  int
}
