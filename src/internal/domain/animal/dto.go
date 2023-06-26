package animal

type Animal struct {
	Id        int
	Name      string
	Age       int
	Height    float64
	Weight    float64
	ShelterId int
	Type      string
	Gender    string
}

type Animals []*Animal

type CreateAnimal struct {
	Name      string
	Age       int
	Height    float64
	Weight    float64
	ShelterId int
	Type      string
	Gender    string
}

type UpdateAnimal struct {
	Name      string
	Age       int
	Height    float64
	Weight    float64
	ShelterId int
	Type      string
	Gender    string
}
