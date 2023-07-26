package duck

import "fmt"

// display
func NormalDisplay() {
	fmt.Println("looks like this is a duck")
}

func RubberDisplay() {
	fmt.Println("this is a rubber duck")
}

func InjuredDisplay() {
	fmt.Println("this duck has a broken wing :(")
}

// quacks
func NormalQuack() {
	fmt.Println("quack!")
}

func MuteQuack() {
	fmt.Println("...")
}

// flies
func NormalFly() {
	fmt.Println("it flies to the next pond over")
}

func NotMoving() {
	fmt.Println("it just stays there")
}

func AtLeastItTried() {
	fmt.Println("it jumps... I guess?")
}
