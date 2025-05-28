package enemy

import (
	"5abstractfactory/product/enemy/interfaces"
	"5abstractfactory/product/model"
	wInterface "5abstractfactory/product/weapon/interfaces"
	"fmt"
	"math/rand"
)

type Lampmaster struct {
	name    string
	stat    model.Stat
	weapons []wInterface.WeaponInterface
}

func NewLampmaster(weapons ...wInterface.WeaponInterface) interfaces.EnemyInterface {
	return &Lampmaster{
		name: "Lampmaster",
		stat: model.Stat{
			Might:      800,
			CritRate:   0.4,
			CritDamage: 1.5,
		},
		weapons: weapons,
	}
}

func (e *Lampmaster) GetProfile() string {

	var minMight, maxMight int
	var minCritRate, maxCritRate, minCritDamage, maxCritDamage float32

	for _, weapon := range e.weapons {
		modifier := weapon.GetModifier()
		might := e.stat.Might + modifier.Might
		critRate := e.stat.CritRate + modifier.CritRate
		critDamage := e.stat.CritDamage + modifier.CritDamage

		if minMight == 0 || might <= minMight {
			minMight = might
		}
		if maxMight == 0 || might >= maxMight {
			maxMight = might
		}
		if minCritRate == 0 || critRate <= minCritRate {
			minCritRate = critRate
		}
		if maxCritRate == 0 || critRate >= maxCritRate {
			maxCritRate = critRate
		}
		if minCritDamage == 0 || critDamage <= minCritDamage {
			minCritDamage = critDamage
		}
		if maxCritDamage == 0 || critDamage >= maxCritDamage {
			maxCritDamage = critDamage
		}
	}

	txtMight := fmt.Sprintf("%d - %d", minMight, maxMight)
	txtCritRate := fmt.Sprintf("%.0f%% - %.0f%%", (minCritRate * 100), (maxCritRate * 100))
	txtCritDamage := fmt.Sprintf("%.0f%% - %.0f%%", (minCritDamage * 100), (maxCritDamage * 100))

	txtStats := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n\n%s",
		fmt.Sprintf("You are facing %s", e.name),
		fmt.Sprintf("Might: %s", txtMight),
		fmt.Sprintf("Crit Rate: %s", txtCritRate),
		fmt.Sprintf("Crit Damage: %s", txtCritDamage),
		"======================================================",
	)

	txtAbilities := "Abilities:\n"
	for _, weapon := range e.weapons {
		abilities := weapon.GetAbilities()

		for _, ability := range abilities {
			txtAbilities += fmt.Sprintf("- %s: %s \n", ability.Name, ability.Description)
		}
	}

	return fmt.Sprintf("%s\n\n%s", txtStats, txtAbilities)
}

func (e *Lampmaster) Attack() string {

	weaponIndex := rand.Intn(len(e.weapons))

	abilities := e.weapons[weaponIndex].GetAbilities()
	abilityIndex := rand.Intn(len(abilities))

	ability := abilities[abilityIndex]

	damage := e.weapons[weaponIndex].CalculateDamage(abilityIndex, e.stat)

	message := fmt.Sprintf("%s attack you with %s and deal a total of %.0f damage.", e.name, ability.Name, damage)

	switch ability.Name {
	case "Disrupt Aim":
		if 0.5 <= rand.Float32() {
			message = fmt.Sprintf("%s You are now dizzy.", message)
		}
	case "Dark Explosion":
		message = fmt.Sprintf("%s You are now exhausted.", message)
	}

	return message

}
