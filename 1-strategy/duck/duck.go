package duck

type Duck struct {
	QuackBehaviour   func()
	FlyBehaviour     func()
	DisplayBehaviour func()
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
