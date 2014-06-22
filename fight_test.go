package main

import (
	"fmt"
	"testing"
)

func TestTakeDamage(t *testing.T) {
	test := armorEquip{
		equipment: equipment{durability: 100},
		armor:     chestplate,
		strength:  80,
		hardness:  75,
		dampening: 50,
	}
	chest := bodyPartInstance{
		armor:    &test,
		bodyPart: chest,
		health:   100,
	}

	fmt.Println(test)
	fmt.Println(chest)

	count := 0

	for chest.health > 0 {
		chest.takeAttack(75, 75, 15)
		fmt.Println(test)
		fmt.Println(chest)
		count++
	}

	fmt.Println(count)

}

func TestAttackBodyPart(t *testing.T) {
	attacker := randPerson()
	defender := randPerson()

	armor := armorEquip{
		equipment: equipment{durability: 100},
		armor:     chestplate,
		strength:  80,
		hardness:  75,
		dampening: 50,
	}

	weapon := weaponEquip{
		equipment: equipment{
			durability: 100,
			owner:      &attacker,
		},
		weapon:    sword,
		strength:  75,
		minDamage: 15,
		maxDamage: 30,
		sharpness: 50,
		bluntness: 25,
	}

	defender.getBodyPart(chest).armor = &armor
	attacker.equipped.weapon = weapon

	for defender.getBodyPart(chest).health > 0 {
		attacker.attackBodyPart(&defender, chest)
		fmt.Println(defender.getBodyPart(chest))

	}

}

func TestGetTarget(t *testing.T) {
	attacker := randPerson()
	defender := randPerson()

	target := attacker.chooseTarget(defender)
	fmt.Println(target)
}

func TestAttack(t *testing.T) {
	attacker := randPerson()
	defender := randPerson()

	attacker.name = "Cody"
	defender.name = "Quint"

	for defender.health.unconscious == false {
		attacker.attack(&defender)
		fmt.Println(defender.health)
		fmt.Println(attacker.rollDamage())
	}
}
