package enemy

import (
	"4factorymethod/product/enemy/interfaces"
	"4factorymethod/product/model"
	"fmt"
	"math/rand"
)

type Lancelier struct {
	stats     model.EnemyStat
	abilities []model.Ability
}

func NewLancelier(stats model.EnemyStat, abilities []model.Ability) interfaces.EnemyInterface {
	return &Lancelier{
		stats:     stats,
		abilities: abilities,
	}
}

func (e *Lancelier) GetStats() model.EnemyStat {
	return e.stats
}

func (e *Lancelier) GetAbilities() []model.Ability {
	return e.abilities
}

func (e *Lancelier) Attack() string {

	randomIndex := rand.Intn(len(e.abilities))
	ability := e.abilities[randomIndex]

	damage := e.stats.Might * ability.AttackPower
	if e.stats.CritRate <= rand.Float32() {
		damage = damage * e.stats.CritDamage
	}

	return fmt.Sprintf("%s attack you with %s and deal %.0f damage.", e.stats.Name, ability.Name, damage)

}
