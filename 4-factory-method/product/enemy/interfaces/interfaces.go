package interfaces

import (
	"4factorymethod/product/model"
)

type EnemyInterface interface {
	// Initiating() EnemyInterface
	GetStats() model.EnemyStat
	GetAbilities() []model.Ability
	Attack() string
}
