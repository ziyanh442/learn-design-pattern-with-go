package enemy

import (
	"4factorymethod/product/enemy/interfaces"
	"4factorymethod/product/model"
	"fmt"
	"math/rand"
)

type Lampmaster struct {
	stats     model.EnemyStat
	abilities []model.Ability
}

func NewLampmaster() interfaces.EnemyInterface {
	return &Lampmaster{
		stats: model.EnemyStat{
			Name:       "Lampmaster",
			Health:     10000,
			Might:      5000,
			CritRate:   0.7,
			CritDamage: 2,
		},
		abilities: []model.Ability{
			{
				Name:        "Jump Attack",
				Description: "Strikes a character once. Deals physical damage.",
				AttackPower: 0.3,
			},
			{
				Name:        "Lamp Attack",
				Description: "Each active Lamp shoots a projectile at a random character. Deals light damage.",
				AttackPower: 0.7,
			},
			{
				Name:        "Disrupt Aim",
				Description: "Strikes a character three times with its sword. Deals physical damage. Can apply Dizzy.",
				AttackPower: 0.5,
			},
			{
				Name:        "Dark Explosion",
				Description: "Strikes a character once. Deals high dark damage. Applies exhaust.",
				AttackPower: 1,
			},
		},
	}
}

func (e *Lampmaster) GetStats() model.EnemyStat {
	return e.stats
}

func (e *Lampmaster) GetAbilities() []model.Ability {
	return e.abilities
}

func (e *Lampmaster) Attack() string {

	randomIndex := rand.Intn(len(e.abilities))
	ability := e.abilities[randomIndex]

	damage := e.stats.Might * ability.AttackPower
	if e.stats.CritRate <= rand.Float32() {
		damage = damage * e.stats.CritDamage
	}

	message := fmt.Sprintf("%s attack you with %s and deal %.0f damage.", e.stats.Name, ability.Name, damage)

	switch ability.Name {
	case "Disrupt Aim":
		if 0.5 <= rand.Float32() {
			message = fmt.Sprintf("%s You are now dizzy.", message)
		}
	case "Dark Explosion":
		message = fmt.Sprintf("%s You are now exhausted.", message)
	}

	return message

}
