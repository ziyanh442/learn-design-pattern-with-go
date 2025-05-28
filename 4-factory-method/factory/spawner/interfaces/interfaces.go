package interfaces

import "4factorymethod/product/enemy/interfaces"

type FactoryInterface interface {
	SpawnEnemy() interfaces.EnemyInterface
}
