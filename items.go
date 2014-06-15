package main

type material int

const (
	wood = iota
	stone
	tin
	copper
	iron
	steel
	adamantine
)

type weapon int

const (
	knife = iota
	sword
	battleAxe
	warHammer
	pike
	bow
	crossbow
	numWeapons
)

type tool int

const (
	hoe = iota
	shovel
	hammer
	axe
	saw
	numTools
)

type armor int

const (
	helmet = iota
	chestplate
	gauntlets
	greaves
	boots
	shield
	numArmor
)

type equipment struct {
	name       string
	quality    int
	durability int
	weight     int
	material   material
	owner      *person
}

type carriable interface {
	getOwner() *person
	getWeight() int
}

type ownable interface {
	getOwner() *person
}

func (e *equipment) getOwner() *person {
	return e.owner
}

func (e *equipment) getWeight() int {
	return e.weight
}

type toolEquip struct {
	equipment
	tool tool
}

type weaponEquip struct {
	equipment
	weapon weapon
}

type armorEquip struct {
	equipment
	armor armor
}
