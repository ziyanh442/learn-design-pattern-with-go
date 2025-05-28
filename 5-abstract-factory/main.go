package main

import (
	"5abstractfactory/factory/spawner"
	"5abstractfactory/factory/spawner/interfaces"
	"fmt"
	"math/rand"
)

func main() {

	var factory interfaces.FactoryInterface

	switch rand.Intn(3) {
	case 0:
		factory = spawner.NewLancelierSpawner()
	case 1:
		factory = spawner.NewBrulerSpawner()
	case 2:
		factory = spawner.NewLampmasterSpawner()
	default:
		panic("out of range")
	}

	enemy := factory.SpawnEnemy(factory.GetWeapons()...)

	fmt.Print("======================================================\n\n")

	fmt.Println(enemy.GetProfile())

	fmt.Print("======================================================\n\n")

	fmt.Println(enemy.Attack())
}
