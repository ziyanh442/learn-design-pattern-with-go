package model

type Stat struct {
	Might      int
	CritRate   float32
	CritDamage float32
}

type Ability struct {
	Name        string
	Description string
	AttackPower float32
}
