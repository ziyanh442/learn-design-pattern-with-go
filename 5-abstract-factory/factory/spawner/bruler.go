package spawner

import (
	"5abstractfactory/factory/interfaces"
	"5abstractfactory/product/enemy"
	eInterface "5abstractfactory/product/enemy/interfaces"
	"5abstractfactory/product/weapon"
	wInterface "5abstractfactory/product/weapon/interfaces"
)

type BrulerSpawner struct{}

func NewBrulerSpawner() interfaces.FactoryInterface {
	return &BrulerSpawner{}

}

func (f *BrulerSpawner) SpawnEnemy(weapons ...wInterface.WeaponInterface) eInterface.EnemyInterface {
	return enemy.NewBruler(weapons[0])
}

func (f *BrulerSpawner) GetWeapons() []wInterface.WeaponInterface {
	return []wInterface.WeaponInterface{
		weapon.NewAnchorWeapon(),
	}
}
