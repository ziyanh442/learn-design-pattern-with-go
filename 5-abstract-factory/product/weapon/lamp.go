package weapon

import (
	"5abstractfactory/product/model"
	"5abstractfactory/product/weapon/interfaces"
	"math/rand"
)

type Lamp struct {
	name     string
	modifier model.Stat
	ability  model.Ability
}

func NewLampWeapon() interfaces.WeaponInterface {
	return &Lamp{
		name: "Lamp",
		modifier: model.Stat{
			Might:      350,
			CritRate:   0.1,
			CritDamage: 0,
		},
		ability: model.Ability{
			Name:        "Lamp Attack",
			Description: "Each active Lamp shoots a projectile at a random character. Deals light damage.",
			AttackPower: 0.7,
		},
	}
}

func (w *Lamp) GetName() string {
	return w.name
}

func (w *Lamp) GetAbilities() []model.Ability {
	return []model.Ability{
		w.ability,
	}
}

func (w *Lamp) GetModifier() model.Stat {
	return w.modifier
}

func (w *Lamp) CalculateDamage(_ int, baseStat model.Stat) (damage float32) {

	mightWithModifier := baseStat.Might + w.modifier.Might
	critRateWithModifier := baseStat.CritRate + w.modifier.CritRate
	critDamageWithModifier := baseStat.CritDamage + w.modifier.CritDamage

	damage = w.ability.AttackPower * float32(mightWithModifier)

	if critRateWithModifier <= rand.Float32() {
		damage = damage * critDamageWithModifier
	}

	return damage
}
