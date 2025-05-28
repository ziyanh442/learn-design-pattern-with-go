package spawner

import (
	"5abstractfactory/factory/spawner/interfaces"
	"5abstractfactory/product/enemy"
	eInterface "5abstractfactory/product/enemy/interfaces"
	"5abstractfactory/product/weapon"
	wInterface "5abstractfactory/product/weapon/interfaces"
)

type LancelierSpawner struct{}

func NewLancelierSpawner() interfaces.FactoryInterface {
	return &LancelierSpawner{}

}

func (f *LancelierSpawner) SpawnEnemy(weapons ...wInterface.WeaponInterface) eInterface.EnemyInterface {
	return enemy.NewLancelier(weapons[0])
}

func (f *LancelierSpawner) GetWeapons() []wInterface.WeaponInterface {
	return []wInterface.WeaponInterface{
		weapon.NewLanceWeapon(),
	}
}
