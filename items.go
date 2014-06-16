package main

// material is the types of materials equipable items can be made of.
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

func (m *material) getStrenght() int {
	switch *m {
	case wood:
		return 2
	case stone:
		return 4
	case tin:
		return 5
	case copper:
		return 6
	case iron:
		return 7
	case steel:
		return 8
	case adamantine:
		return 10
	default:
		return 0
	}
}

// weapon denotes the types of weapons.
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

// tool denotes the types of tools.
type tool int

const (
	hoe = iota
	shovel
	hammer
	axe
	saw
	numTools
)

// armor denotes the types of armor.
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

// equipment struct represents an instance of an item.
type equipment struct {
	name       string
	quality    int
	durability int
	weight     int
	material   material
	owner      *person
}

// carriable interface means a character can carry an item.
type carriable interface {
	getOwner() *person
	getWeight() int
}

// ownable interface allows a character to own something
type ownable interface {
	getOwner() *person
}

// getOwner returns the owner of equipment.
func (e *equipment) getOwner() *person {
	return e.owner
}

// getWeight returns the weight of equipment.
func (e *equipment) getWeight() int {
	return e.weight
}

// toolEquip is a tool instance.
type toolEquip struct {
	equipment
	tool tool
}

// weaonEquip is a weapon instance.
type weaponEquip struct {
	equipment
	weapon    weapon
	minDamage int
	maxDamage int
	armorPen  int
}

// armorEquip is an armor instance.
type armorEquip struct {
	equipment
	armor armor
}
