package main

import (
	"1strategy/duck"
	"fmt"
)

func main() {
	normalDuck := duck.NewDuck(duck.NormalQuack, duck.NormalFly, duck.NormalDisplay)
	normalDuck.Display()
	normalDuck.Quack()
	normalDuck.Fly()

	fmt.Println("======================")

	rubberDuck := duck.NewDuck(duck.MuteQuack, duck.NotMoving, duck.RubberDisplay)
	rubberDuck.Display()
	rubberDuck.Quack()
	rubberDuck.Fly()

	fmt.Println("======================")

	injuredDuck := duck.NewDuck(duck.NormalQuack, duck.AtLeastItTried, duck.InjuredDisplay)
	injuredDuck.Display()
	injuredDuck.Quack()
	injuredDuck.Fly()
}
