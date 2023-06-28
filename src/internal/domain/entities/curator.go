package entities

type Curator struct {
	Id          int
	ChatId      string
	Name        string
	Surname     string
	PhoneNumber string
}

type Curators []*Curator
