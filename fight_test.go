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
		dampening: 50,
	}

	fmt.Println(test)
	fmt.Println(chest)

	count := 0

	for chest.health > 0 {
		test.takeDamage(75, 75, 15)
		fmt.Println(test)
		fmt.Println(chest)
		count++
	}

	fmt.Println(count)

}
