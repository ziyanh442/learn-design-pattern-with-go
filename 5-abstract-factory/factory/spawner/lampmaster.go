package spawner

import (
	"5abstractfactory/factory/interfaces"
	"5abstractfactory/product/enemy"
	eInterface "5abstractfactory/product/enemy/interfaces"
	"5abstractfactory/product/weapon"
	wInterface "5abstractfactory/product/weapon/interfaces"
)

type LampmasterSpawner struct{}

func NewLampmasterSpawner() interfaces.FactoryInterface {
	return &LampmasterSpawner{}

}

func (f *LampmasterSpawner) SpawnEnemy(weapons ...wInterface.WeaponInterface) eInterface.EnemyInterface {
	return enemy.NewLampmaster(weapons...)
}

func (f *LampmasterSpawner) GetWeapons() []wInterface.WeaponInterface {
	return []wInterface.WeaponInterface{
		weapon.NewLampWeapon(),
		weapon.NewSwordWeapon(),
	}
}
