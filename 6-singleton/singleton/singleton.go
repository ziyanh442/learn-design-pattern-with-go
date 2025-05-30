package singleton

import (
	"fmt"
	"math/rand"
)

var db dbInterface

func GetDatabase() dbInterface {
	if db != nil {
		fmt.Println("returning existing database instance...")
		return db
	}

	fmt.Println("creating a new database instance...")
	db = newDatabaseInstance(rand.Intn(100))

	return db
}
