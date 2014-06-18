package main

import (
	"fmt"
	"testing"
)

func TestTakeDamage(t *testing.T) {
	chest := bodyPartInstance{
		bodyPart: chest,
		health:   100,
	}
	test := armorEquip{
		equipment: equipment{durability: 100},
		armor:     chestplate,
		equipedOn: &chest,
		strength:  80,
		hardness:  75,
		dampening: 75,
	}

	fmt.Println(test)
	fmt.Println(chest)

	for chest.health > 0 {
		test.takeDamage(30, 0, 0)
		fmt.Println(test)
		fmt.Println(chest)
	}

}
