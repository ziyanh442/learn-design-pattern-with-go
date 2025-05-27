package model

type EnemyStat struct {
	Name       string
	Health     float32
	Might      float32
	CritRate   float32
	CritDamage float32
}

type Ability struct {
	Name        string
	Description string
	AttackPower float32
}
