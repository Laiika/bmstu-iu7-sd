package curator

type Curator struct {
	Id          int
	ChatId      string
	Name        string
	Surname     string
	PhoneNumber string
}

type Curators []*Curator

type CreateCurator struct {
	ChatId      string
	Name        string
	Surname     string
	PhoneNumber string
}

type UpdateCurator struct {
	ChatId      string
	Name        string
	Surname     string
	PhoneNumber string
}
