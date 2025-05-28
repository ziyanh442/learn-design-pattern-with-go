package weapon

import (
	"5abstractfactory/product/model"
	"5abstractfactory/product/weapon/interfaces"
	"math/rand"
)

type Lance struct {
	name      string
	modifier  model.Stat
	abilities []model.Ability
}

func NewLanceWeapon() interfaces.WeaponInterface {
	return &Lance{
		name: "Lance",
		modifier: model.Stat{
			Might:      200,
			CritRate:   0.3,
			CritDamage: 0.3,
		},
		abilities: []model.Ability{
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
		},
	}
}

func (w *Lance) GetName() string {
	return w.name
}

func (w *Lance) GetAbilities() []model.Ability {
	return w.abilities
}

func (w *Lance) GetModifier() model.Stat {
	return w.modifier
}

func (w *Lance) CalculateDamage(abilityIndex int, baseStat model.Stat) (damage float32) {

	mightWithModifier := baseStat.Might + w.modifier.Might
	critRateWithModifier := baseStat.CritRate + w.modifier.CritRate
	critDamageWithModifier := baseStat.CritDamage + w.modifier.CritDamage

	damage = w.abilities[abilityIndex].AttackPower * float32(mightWithModifier)

	if critRateWithModifier <= rand.Float32() {
		damage = damage * critDamageWithModifier
	}

	return damage
}
