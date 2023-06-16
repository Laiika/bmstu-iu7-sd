package animal

type Animal struct {
	Id        int32
	Name      string
	Age       int32
	Height    int32
	Weight    float32
	ShelterId int32
	Type      string
	Sex       string
}

type Animals []*Animal

type CreateAnimal struct {
	Name      string
	Age       int32
	Height    int32
	Weight    float32
	ShelterId int32
	Type      string
	Sex       string
}

type UpdateAnimal struct {
	Name      string
	Age       int32
	Height    int32
	Weight    float32
	ShelterId int32
	Type      string
	Sex       string
}
