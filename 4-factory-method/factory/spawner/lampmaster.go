package spawner

import (
	"4factorymethod/factory/spawner/interfaces"
	"4factorymethod/product/enemy"
	enemyInterface "4factorymethod/product/enemy/interfaces"
)

type LampmasterSpawner struct{}

func NewLampmasterSpawner() interfaces.FactoryInterface {
	return &LampmasterSpawner{}
}

func (s *LampmasterSpawner) SpawnEnemy() enemyInterface.EnemyInterface {
	return enemy.NewLampmaster()
}
