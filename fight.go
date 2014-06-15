package main

import "math/rand"

func Fight(x *Person, y *Person) {

}

//TODO factor in weight load
func rollInitiative(p Person, stamina int) int {
	initiative := p.GetWeaponSkill() + stamina + p.Attributes.Agility
	return int(rand.Int31n(int32(initiative)))
}

func calcCombatSkill(p Person) int {
	sum := 0
	for x := 0; x < numWeapons; x++ {
		sum += p.Skills.Weapons[Weapon(x)]
	}
	for x := 0; x < numArmor; x++ {
		sum += p.Skills.ArmorUse[Armor(x)]
	}

	return sum / (numWeapons + numArmor)

}
