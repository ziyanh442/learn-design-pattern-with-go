package main

import (
	"4factorymethod/factory/interfaces"
	"4factorymethod/factory/spawner"
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

	enemy := factory.SpawnEnemy()
	stats := enemy.GetStats()
	abilities := enemy.GetAbilities()

	fmt.Println("======================================================")

	fmt.Printf("You are facing %s \n", stats.Name)
	fmt.Printf("Might: %.0f \n", stats.Might)
	fmt.Printf("Critical Rate: %.0f%% \n", (stats.CritRate * 100))
	fmt.Printf("Critical Damage: %.0f%% \n", (stats.CritDamage * 100))

	fmt.Println("======================================================")

	fmt.Println("Abilities:")
	for _, ability := range abilities {
		fmt.Printf("%s: %s \n", ability.Name, ability.Description)
	}
	fmt.Print("======================================================\n\n")

	fmt.Println(enemy.Attack())
}
