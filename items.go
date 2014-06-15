package main

type Material int

const (
	Wood = iota
	Stone
	Tin
	Copper
	Iron
	Steel
	Adamantine
)

type Weapon int

const (
	Knife = iota
	Sword
	BattleAxe
	WarHammer
	Pike
	Bow
	Crossbow
	numWeapons
)

type Tool int

const (
	Hoe = iota
	Shovel
	Hammer
	Axe
	Saw
	numTools
)

type Armor int

const (
	Helmet = iota
	Chestplate
	Gauntlets
	Greaves
	Boots
	Shield
	numArmor
)

type Equipment struct {
	Name       string
	Quality    int
	Durability int
	Weight     int
	Material   Material
	Owner      *Person
}

type Carriable interface {
	GetOwner() *Person
	GetWeight() int
}

type Ownable interface {
	GetOwner() *Person
}

func (e *Equipment) GetOwner() *Person {
	return e.Owner
}

func (e *Equipment) GetWeight() int {
	return e.Weight
}

type ToolEquip struct {
	Equipment
	Tool Tool
}

type WeaponEquip struct {
	Equipment
	Weapon Weapon
}

type ArmorEquip struct {
	Equipment
	Armor Armor
}
