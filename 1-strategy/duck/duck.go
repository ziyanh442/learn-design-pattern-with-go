package duck

type DuckInterface interface {
	Quack()
	Fly()
	Display()
}

type Duck struct {
	QuackBehaviour   func()
	FlyBehaviour     func()
	DisplayBehaviour func()
}

func NewDuck(quackBehaviour func(), flyBehaviour func(), displayBehaviour func()) DuckInterface {
	return &Duck{
		QuackBehaviour:   quackBehaviour,
		FlyBehaviour:     flyBehaviour,
		DisplayBehaviour: displayBehaviour,
	}
}

func (d *Duck) Quack() {
	d.QuackBehaviour()
}

func (d *Duck) Fly() {
	d.FlyBehaviour()
}

func (d *Duck) Display() {
	d.DisplayBehaviour()
}
