package main

import (
	"6singleton/singleton"
	"fmt"
)

func main() {

	db1 := singleton.GetDatabase()
	db2 := singleton.GetDatabase()

	fmt.Println("id db1: ", db1.GetID())
	fmt.Println("id db2: ", db2.GetID())

	fmt.Println(db2.Select())
	fmt.Println(db2.Update())
	fmt.Println(db2.Delete())

	db3 := singleton.GetDatabase()

	fmt.Println("id db3: ", db3.GetID())
	fmt.Println(db3.Select())
	fmt.Println(db3.Update())
	fmt.Println(db3.Delete())

}
