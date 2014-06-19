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

func (b *bodyPart) getAccNeeded() int {
	switch *b {
	case head:
		return 70
	case lArm:
		return 80
	case rArm:
		return 80
	case lLeg:
		return 80
	case rLeg:
		return 80
	case chest:
		return 50
	default:
		return 90

	}
}

func (p *person) attack(t *person, bp bodyPart) {
	accuracy := p.rollAccuracy(t, bp)

	var contact int

	if accuracy >= 0 {
		if accuracy >= 20 {
			contact = 100
		} else {
			contact = int(float32(accuracy)/20*100 + .5)
		}
	} else {
		if accuracy >= -20 {
			bp = bodyPart(rand.Int31n(numBodyParts - 1))
			contact = int(rand.Int31n(50))
		} else {
			contact = 0
		}

	}

	contact += 0
	// TODO left off here, need to do the damage

}

func (p *person) rollAccuracy(t *person, bp bodyPart) int {

	accuracy := p.getWeaponSkill() + p.accuracy + p.agility

	// apply target dodge attempt
	dodge := calcDamageRatio(t.agility,
		intRoundDiv((int(t.movementCapacity())+t.agility), 2))

	if dodge >= accuracy {
		accuracy = 0
	} else {
		accuracy -= dodge
	}

	accuracy = calcDamageRatio(accuracy,
		intRoundDiv(p.health.stamina+int(p.movementCapacity()), 2))

	need := bp.getAccNeeded()

	need -= int(float32(accuracy)/30*50 + .5)

	roll := rand.Int31n(100)

	return int(roll) - need

}

func (b *bodyPartInstance) takeDamage(blunt int, cut int) {
	if blunt >= 30 {
		b.broken = true
	}
	if cut >= 30 {
		b.severed = true
	}
	if cut+blunt >= 50 {
		b.detached = true
	}
	if b.health <= cut+blunt {
		b.health = 0
	} else {
		b.health -= cut + blunt
	}
}

func (a *armorEquip) takeDamage(damage int, blunt int, cut int) {

	bluntTransfered := calcDamageRatio(damage, blunt)
	cutTransfered := calcDamageRatio(damage, cut)

	bluntDamage := bluntTransfered - calcDamageRatio(bluntTransfered, a.getDampening())
	cutDamage := cutTransfered - calcDamageRatio(cutTransfered, a.getHardness())

	a.equipedOn.takeDamage(bluntDamage, cutDamage)

	armorDamage := damage - calcDamageRatio(damage, a.strength)

	if armorDamage >= a.durability {
		a.durability = 0
	} else {
		a.durability -= armorDamage
	}
}

func calcDamageRatio(damage int, ratio int) int {

	damage = int(float32(damage)*(float32(ratio)/100.0) + .5)

	return damage
}
