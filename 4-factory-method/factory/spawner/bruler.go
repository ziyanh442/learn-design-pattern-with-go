package spawner

import (
	"4factorymethod/factory/interfaces"
	"4factorymethod/product/enemy"
	enemyInterface "4factorymethod/product/interfaces"
)

type BrulerSpawner struct{}

func NewBrulerSpawner() interfaces.FactoryInterface {
	return &BrulerSpawner{}
}

func (s *BrulerSpawner) SpawnEnemy() enemyInterface.EnemyInterface {
	return enemy.NewBruler()
}

//it may be a better example if each spawner has something unique going on instead of simply calling the constructor method
