package singleton

type dbInterface interface {
	GetID() int
	Select() string
	Update() string
	Delete() string
}

type dbConcreteImpl struct {
	id int //store a randomized integer just to illustrate we are still calling the same instance
}

func newDatabaseInstance(id int) dbInterface {
	return &dbConcreteImpl{
		id: id,
	}
}

func (d *dbConcreteImpl) GetID() int {
	return d.id
}

func (d *dbConcreteImpl) Select() string {
	return "calling Select method"
}

func (d *dbConcreteImpl) Update() string {
	return "calling Update method"
}

func (d *dbConcreteImpl) Delete() string {
	return "calling Delete method"
}
