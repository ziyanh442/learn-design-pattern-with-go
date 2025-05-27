package spawner

import (
	"4factorymethod/factory/interfaces"
	"4factorymethod/product/enemy"
	enemyInterface "4factorymethod/product/interfaces"
	"4factorymethod/product/model"
)

type LancelierSpawner struct{}

func NewLancelierSpawner() interfaces.FactoryInterface {
	return &LancelierSpawner{}
}

func (s *LancelierSpawner) SpawnEnemy() enemyInterface.EnemyInterface {
	return enemy.NewLancelier(model.EnemyStat{
		Name:       "Lancelier",
		Health:     5000,
		Might:      300,
		CritRate:   0.4,
		CritDamage: 1.5,
	}, []model.Ability{
		{
			Name:        "Quick Attack",
			Description: "A quick single strike that deals low physical damage.",
			AttackPower: 0.2,
		},
		{
			Name:        "Slow Attack",
			Description: "A slow wind-up strike that deals moderate physical damage, that causes the Lancelier's weapon to glow red before attacking.",
			AttackPower: 0.5,
		},
	})
}
