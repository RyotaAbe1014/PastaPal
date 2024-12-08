package recipes

type Unit struct {
	id          int32
	name        string
	description string
}

func NewUnit(id int32, name string, description string) Unit {
	return Unit{id: id, name: name, description: description}
}

func (u Unit) ID() int32 {
	return u.id
}

func (u Unit) Name() string {
	return u.name
}

func (u Unit) Description() string {
	return u.description
}
