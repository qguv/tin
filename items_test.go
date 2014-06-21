package main

import "testing"

func TestArmorEquip(t *testing.T) {
	test := armorEquip{
		equipment: equipment{durability: 100},
		armor:     chestplate,
		strength:  90,
		hardness:  75,
		dampening: 75,
	}

	if test.getDampening() != 75 {
		t.Fail()
	}
	if test.getHardness() != 75 {
		t.Fail()
	}

	test.durability = 50

	if test.getDampening() != 38 || test.getHardness() != 38 {
		t.Fail()
	}
}
