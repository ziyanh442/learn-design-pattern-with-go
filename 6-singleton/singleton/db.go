package singleton

type dbInterface interface {
	GetID() int
	Select() string
	Update() string
	Delete() string
}

type dbClass struct {
	id int //store a randomized integer just to illustrate we are still calling the same instance
}

func newDatabaseInstance(id int) dbInterface {
	return &dbClass{
		id: id,
	}
}

func (d *dbClass) GetID() int {
	return d.id
}

func (d *dbClass) Select() string {
	return "calling Select method"
}

func (d *dbClass) Update() string {
	return "calling Update method"
}

func (d *dbClass) Delete() string {
	return "calling Delete method"
}
