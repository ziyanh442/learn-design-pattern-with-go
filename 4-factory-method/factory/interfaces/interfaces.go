package interfaces

import "4factorymethod/product/interfaces"

type FactoryInterface interface {
	SpawnEnemy() interfaces.EnemyInterface
}
