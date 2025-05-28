package weapon

import (
	"5abstractfactory/product/model"
	"5abstractfactory/product/weapon/interfaces"
	"math/rand"
)

type Anchor struct {
	name     string
	modifier model.Stat
	ability  model.Ability
}

func NewAnchorWeapon() interfaces.WeaponInterface {
	return &Anchor{
		name: "Anchor",
		modifier: model.Stat{
			Might:      300,
			CritRate:   0.1,
			CritDamage: 0,
		},
		ability: model.Ability{
			Name:        "Anchor Attack",
			Description: "Will swing twice at a character. Deals physical damage.",
			AttackPower: 0.6,
		},
	}
}

func (w *Anchor) GetName() string {
	return w.name
}

func (w *Anchor) GetAbilities() []model.Ability {
	return []model.Ability{
		w.ability,
	}
}

func (w *Anchor) GetModifier() model.Stat {
	return w.modifier
}

func (w *Anchor) CalculateDamage(_ int, baseStat model.Stat) (damage float32) {

	mightWithModifier := baseStat.Might + w.modifier.Might
	critRateWithModifier := baseStat.CritRate + w.modifier.CritRate
	critDamageWithModifier := baseStat.CritDamage + w.modifier.CritDamage

	damage = w.ability.AttackPower * float32(mightWithModifier)

	if critRateWithModifier <= rand.Float32() {
		damage = damage * critDamageWithModifier
	}

	return damage
}
