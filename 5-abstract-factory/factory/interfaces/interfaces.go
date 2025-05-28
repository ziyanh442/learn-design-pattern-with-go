package interfaces

import (
	eInterface "5abstractfactory/product/enemy/interfaces"
	wInterface "5abstractfactory/product/weapon/interfaces"
)

type FactoryInterface interface {
	SpawnEnemy(...wInterface.WeaponInterface) eInterface.EnemyInterface
	GetWeapons() []wInterface.WeaponInterface
}
