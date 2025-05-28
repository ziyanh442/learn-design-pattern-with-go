package enemy

import (
	"5abstractfactory/product/enemy/interfaces"
	"5abstractfactory/product/model"
	wInterface "5abstractfactory/product/weapon/interfaces"
	"fmt"
)

type Bruler struct {
	name   string
	stat   model.Stat
	weapon wInterface.WeaponInterface
}

func NewBruler(weapon wInterface.WeaponInterface) interfaces.EnemyInterface {
	return &Bruler{
		name: "Bruler",
		stat: model.Stat{
			Might:      500,
			CritRate:   0.2,
			CritDamage: 1.2,
		},
		weapon: weapon,
	}
}

func (e *Bruler) GetProfile() string {

	modifier := e.weapon.GetModifier()
	abilities := e.weapon.GetAbilities()

	critRate := (e.stat.CritRate + modifier.CritRate) * 100
	critDamage := (e.stat.CritDamage + modifier.CritDamage) * 100

	txtStats := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n\n%s",
		fmt.Sprintf("You are facing %s", e.name),
		fmt.Sprintf("Might: %d", (e.stat.Might+modifier.Might)),
		fmt.Sprintf("Crit Rate: %.0f%%", critRate),
		fmt.Sprintf("Crit Damage: %.0f%%", critDamage),
		"======================================================",
	)

	txtAbilities := "Abilities:\n"
	for _, ability := range abilities {
		txtAbilities += fmt.Sprintf("- %s: %s \n", ability.Name, ability.Description)
	}

	return fmt.Sprintf("%s\n\n%s", txtStats, txtAbilities)
}

func (e *Bruler) Attack() string {

	abilityIndex := 0
	ability := e.weapon.GetAbilities()[abilityIndex]

	damage := e.weapon.CalculateDamage(abilityIndex, e.stat)

	return fmt.Sprintf("%s attack you with %s and deal a total of %.0f damage.", e.name, ability.Name, damage)

}
