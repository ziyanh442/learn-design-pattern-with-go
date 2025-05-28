package interfaces

import "5abstractfactory/product/model"

type WeaponInterface interface {
	GetName() string
	GetModifier() model.Stat
	GetAbilities() []model.Ability
	CalculateDamage(abilityIndex int, baseStat model.Stat) (damage float32)
}
