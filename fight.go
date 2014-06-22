package main

import (
	"fmt"
	"math/rand"
)

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

// getAccNeeded returns the base roll needed to hit
// a given target area
func (b *bodyPart) getAccNeeded() int {
	switch *b {
	case head:
		return 60
	case lArm:
		return 70
	case rArm:
		return 70
	case lLeg:
		return 70
	case rLeg:
		return 70
	case chest:
		return 50
	default:
		return 85

	}
}

func (p *person) attack(t *person) {
	target := p.chooseTarget(*t)
	fmt.Printf("%s aims his %s at %s's %s\n", p.name, p.getCurrentWeapon().weapon, t.name, target)
	p.attackBodyPart(t, target)
}

// attack causes a character to attempt to hit the target area
// with current equipped weapon, opponent will attempt to dodge
func (p *person) attackBodyPart(t *person, bp bodyPart) {
	accuracy := p.rollAccuracy(t, bp)

	// fmt.Printf("Accuracy = %d\n", accuracy)
	var contact int

	if accuracy >= 0 {
		if accuracy >= 20 {
			contact = 100
			fmt.Printf("A solid hit!\n")
		} else {
			contact = int(float32(accuracy)/20*100 + .5)
			fmt.Printf("%d contact\n", contact)
		}
	} else {
		if accuracy >= -20 {
			bp = bodyPart(rand.Int31n(numBodyParts - 1))
			contact = int(rand.Int31n(50))
			fmt.Printf("Misses, but hits %s\n", bp)
		} else {
			contact = 0
			fmt.Printf("%s dodges the blow!\n", t.name)
		}

	}

	damage := p.rollDamage()
	// fmt.Printf("Damage = %d\n", damage)

	damage = calcDamageRatio(damage, contact)
	// fmt.Printf("After contact = %d\n", damage)

	target := t.getBodyPart(bp)
	// fmt.Println(target)
	target.takeAttack(damage, p.getWeaponBluntness(), p.getWeaponSharpness())

	// update target health
	t.checkHealth()

	// reduce weapon durability
	p.getCurrentWeapon().reduceDurability(damage, *t.getBodyPart(bp))

}

// rollDamage returns the damage dealt by an attack
func (p *person) rollDamage() int {
	var max int

	max = p.getCurrentWeapon().maxDamage + p.strength + p.getWeaponSkill()

	roll := rand.Int31n(int32(max) - int32(p.getCurrentWeapon().minDamage))

	return int(roll) + p.getCurrentWeapon().minDamage
}

// rollAccuracy returns the differance between roll and needed to hit
// IE, -10 means 10 short of scoring a hit
func (p *person) rollAccuracy(t *person, bp bodyPart) int {

	accuracy := p.getWeaponSkill() + p.accuracy + p.agility

	moveAbility := int(float32(p.health.stamina)*p.movementCapacity() + .5)
	tMoveAbility := int(float32(t.health.stamina)*t.movementCapacity() + .5)

	// apply target dodge attempt
	dodge := t.agility

	need := bp.getAccNeeded()

	need -= int(float32(accuracy)/30*50 + .5)
	need += int(float32(dodge)/10*50 + .5)

	abilityDiff := percentDiff(moveAbility, tMoveAbility)

	if moveAbility > tMoveAbility {
		need -= int(float64(need)*abilityDiff + .5)
	} else {
		need += int(float64(need)*abilityDiff + .5)
	}

	roll := rand.Int31n(100)

	fmt.Printf("acc: %d dodge: %d need: %d roll: %d\n", accuracy, dodge, need, roll)

	return int(roll) - need

}

// takeAttack causes a bodyPartInstance to receive an attack
// if the bodyPart has armor equipped, it is used to block
func (b *bodyPartInstance) takeAttack(damage int, blunt int, cut int) {

	// fmt.Println(b)
	if b.armor != nil {
		// fmt.Println("has armor")
		blunt, cut = b.armor.takeDamage(damage, blunt, cut)
	} else {
		blunt = calcDamageRatio(damage, blunt)
		cut = calcDamageRatio(damage, cut)
	}

	// fmt.Println("taking damage")
	b.takeDamage(blunt, cut)

}

// takeDamage causes a body part to take the indicated amount of damage
// this method applies status effects based on damage
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

// takeDamage calculates the damage that penetrates armor
// it also handles loss of armor durabillity
// TODO handle loss of weapon durabillity
func (a *armorEquip) takeDamage(damage int, blunt int, cut int) (int, int) {

	bluntTransfered := calcDamageRatio(damage, blunt)
	cutTransfered := calcDamageRatio(damage, cut)

	bluntDamage := bluntTransfered - calcDamageRatio(bluntTransfered, a.getDampening())
	cutDamage := cutTransfered - calcDamageRatio(cutTransfered, a.getHardness())

	armorDamage := damage - calcDamageRatio(damage, a.strength)

	if armorDamage >= a.durability {
		a.durability = 0
	} else {
		a.durability -= armorDamage
	}

	return bluntDamage, cutDamage
}

// TODO change this name
func calcDamageRatio(damage int, ratio int) int {

	damage = int(float32(damage)*(float32(ratio)/100.0) + .5)

	return damage
}

func (p *person) chooseTarget(t person) bodyPart {
	var maxDamage int
	var maxTarget bodyPart

	wdamage := p.rollDamage()
	wblunt := p.getWeaponBluntness()
	wcut := p.getWeaponSharpness()

	var blunt, cut int

	for _, target := range t.bodyParts {
		if target.armor != nil {
			blunt, cut = target.armor.takeDamage(wdamage, wblunt, wcut)
		} else {
			blunt = calcDamageRatio(wdamage, wblunt)
			cut = calcDamageRatio(wdamage, wcut)
		}

		damage := blunt + cut

		if target.health-damage <= 0 {
			damage = target.health
		}

		chance := 100 - target.bodyPart.getAccNeeded()

		damage = calcDamageRatio(damage, chance)

		if damage > maxDamage {
			maxDamage = damage
			maxTarget = target.bodyPart
		} else if damage == maxDamage {
			if rand.Int31n(2) == 1 {
				maxTarget = target.bodyPart
			}
		}

	}

	return maxTarget
}
