package main

import "math/rand"

// fight function causes 2 characters to fight.
// this will need to be expanded for group combat.
func fight(x *person, y *person) {

}

// rollInitiative takes in a person and that persons current stamina
// (differant from stamina attribute, which will stay the same)
// and returns a roll for that characters intiative.
// calcualtion based on current stamina, agility, and health
//TODO factor in weight load and health
func rollInitiative(p person, stamina int) int {
	initiative := stamina + p.attributes.agility
	return int(rand.Int31n(int32(initiative)))
}

// calcCombatSkill returns a persons average skill in all combat
// attributes (weapons and armor).
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

type attackable interface {
	defend()
}

func (p *person) hit(t *person, bp bodyPart, contact int) {
}
