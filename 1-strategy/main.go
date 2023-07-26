package main

import (
	"1strategy/duck"
	"fmt"
)

func main() {
	normalDuck := duck.Duck{
		QuackBehaviour:   duck.NormalQuack,
		FlyBehaviour:     duck.NormalFly,
		DisplayBehaviour: duck.NormalDisplay,
	}
	normalDuck.Display()
	normalDuck.Quack()
	normalDuck.Fly()

	fmt.Println("==================")

	rubberDuck := duck.Duck{
		QuackBehaviour:   duck.MuteQuack,
		FlyBehaviour:     duck.NotMoving,
		DisplayBehaviour: duck.RubberDisplay,
	}
	rubberDuck.Display()
	rubberDuck.Quack()
	rubberDuck.Fly()

	fmt.Println("==================")

	injuredDuck := duck.Duck{
		QuackBehaviour:   duck.NormalQuack,
		FlyBehaviour:     duck.AtLeastItTried,
		DisplayBehaviour: duck.InjuredDisplay,
	}
	injuredDuck.Display()
	injuredDuck.Quack()
	injuredDuck.Fly()
}
