package enemy

import (
	"4factorymethod/product/interfaces"
	"4factorymethod/product/model"
	"fmt"
	"math/rand"
)

type Bruler struct {
	stats     model.EnemyStat
	abilities []model.Ability
}

func NewBruler() interfaces.EnemyInterface {
	return &Bruler{
		stats: model.EnemyStat{
			Name:       "Bruler",
			Health:     8000,
			Might:      500,
			CritRate:   0.5,
			CritDamage: 1.5,
		},
		abilities: []model.Ability{
			{
				Name:        "Anchor Attack",
				Description: "Will swing twice at a character. Deals physical damage.",
				AttackPower: 0.6,
			},
		},
	}
}

func (e *Bruler) GetStats() model.EnemyStat {
	return e.stats
}

func (e *Bruler) GetAbilities() []model.Ability {
	return e.abilities
}

func (e *Bruler) Attack() string {

	ability := e.abilities[0]

	damage := e.stats.Might * ability.AttackPower
	if e.stats.CritRate <= rand.Float32() {
		damage = damage * e.stats.CritDamage
	}

	return fmt.Sprintf("%s attack you with %s and deal a total of %.0f damage.", e.stats.Name, ability.Name, damage)

}
