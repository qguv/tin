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
func (p *person) rollInitiative() int {
	initiative := p.health.stamina + p.attributes.agility
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

func (b *bodyPartInstance) takeDamage(blunt int, cut int) {
	if blunt >= 30 {
		b.broken = true
	}
	if cut >= 30 {
		b.severed = true
	}
	if b.health <= cut+blunt {
		b.health = 0
	} else {
		b.health -= cut + blunt
	}
}

func (a *armorEquip) takeDamage(damage int, blunt int, cut int) {
	half := int(float32(damage)/2 + .5)
	bluntDamage := calcArmorPen(half, blunt, a.getDampening())
	cutDamage := calcArmorPen(half, cut, a.getHardness())
	a.equipedOn.takeDamage(bluntDamage, cutDamage)

	armorDamage := int(float32(damage)*(float32(100-a.strength)/100) + .5)

	if armorDamage >= a.durability {
		a.durability = 0
	} else {
		a.durability -= armorDamage
	}
}

func calcArmorPen(damage int, penetration int, resist int) int {

	var res int
	if penetration >= resist {
		res = 0
	} else {
		res = resist - penetration
	}

	blocked := int(float32(damage)*(float32(res)/100.0) + .5)

	return damage - blocked
}
