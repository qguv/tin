package main

import "math/rand"

func fight(x *person, y *person) {

}

//TODO factor in weight load
func rollInitiative(p person, stamina int) int {
	initiative := p.getWeaponSkill() + stamina + p.attributes.agility
	return int(rand.Int31n(int32(initiative)))
}

func calcCombatSkill(p person) int {
	sum := 0
	for x := 0; x < numWeapons; x++ {
		sum += p.skills.weapons[weapon(x)]
	}
	for x := 0; x < numArmor; x++ {
		sum += p.skills.armorUse[armor(x)]
	}

	return sum / (numWeapons + numArmor)

}
