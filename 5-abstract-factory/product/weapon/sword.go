package weapon

import (
	"5abstractfactory/product/model"
	"5abstractfactory/product/weapon/interfaces"
	"math/rand"
)

type Sword struct {
	name      string
	modifier  model.Stat
	abilities []model.Ability
}

func NewSwordWeapon() interfaces.WeaponInterface {
	return &Sword{
		name: "Sword",
		modifier: model.Stat{
			Might:      500,
			CritRate:   0.5,
			CritDamage: 0.5,
		},
		abilities: []model.Ability{
			{
				Name:        "Jump Attack",
				Description: "Strikes a character once. Deals physical damage.",
				AttackPower: 0.3,
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

func (w *Sword) GetName() string {
	return w.name
}

func (w *Sword) GetAbilities() []model.Ability {
	return w.abilities
}

func (w *Sword) GetModifier() model.Stat {
	return w.modifier
}

func (w *Sword) CalculateDamage(abilityIndex int, baseStat model.Stat) (damage float32) {

	mightWithModifier := baseStat.Might + w.modifier.Might
	critRateWithModifier := baseStat.CritRate + w.modifier.CritRate
	critDamageWithModifier := baseStat.CritDamage + w.modifier.CritDamage

	damage = w.abilities[abilityIndex].AttackPower * float32(mightWithModifier)

	if critRateWithModifier <= rand.Float32() {
		damage = damage * critDamageWithModifier
	}

	return damage
}
