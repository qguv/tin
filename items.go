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

// getStrength returns a materials strenght
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
	fist = iota
	knife
	sword
	battleAxe
	warHammer
	pike
	bow
	crossbow
	numWeapons
)

func (w weapon) String() string {
	switch w {
	case fist:
		return "fist"
	case knife:
		return "knife"
	case sword:
		return "sword"
	case battleAxe:
		return "battle-axe"
	case warHammer:
		return "warhammer"
	case pike:
		return "pike"
	case bow:
		return "bow"
	case crossbow:
		return "crossbow"
	default:
		return ""
	}
}

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

type equipable interface {
	equip(*person)
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

// weaponEquip is a weapon instance.
type weaponEquip struct {
	equipment
	weapon    weapon
	strength  int
	minDamage int
	maxDamage int
	sharpness int
	bluntness int
}

func (a *armorEquip) equip(p *person) {
	p.equipped.armor[a.armor].dequip()
	p.equipped.armor[a.armor] = a

	for x, _ := range p.bodyParts {
		if p.bodyParts[x].bodyPart.armorType() == a.armor {
			p.bodyParts[x].armor = a
		}
	}
}

func (w *weaponEquip) equip(p *person) {
	p.equipped.tool = nil
	p.equipped.weapon = w
}

func (t *toolEquip) equipTool(p *person) {
	p.equipped.weapon = nil
	p.equipped.tool = t
}

func (w *weaponEquip) dequip() {
	w.getOwner().equipped.weapon = nil
}

func (a *armorEquip) dequip() {
	bps := &a.getOwner().health.bodyParts

	for x, bp := range *bps {
		if bp.armor == a {
			a.getOwner().health.bodyParts[x].armor = nil
		}
	}

	a.getOwner().equipped.armor[a.armor] = nil
}

func (t *toolEquip) dequip() {
	t.getOwner().equipped.tool = nil
}

func (w *weaponEquip) reduceDurability(damage int, bp bodyPartInstance) {

	const minRedux int = 5

	var redux int
	armor := bp.armor

	blunt := calcDamageRatio(damage, w.bluntness)
	cut := calcDamageRatio(damage, w.sharpness)

	if armor != nil {
		tblunt, tcut := armor.takeDamage(damage, w.bluntness, w.sharpness)
		blocked := blunt - tblunt + cut - tcut
		redux = calcDamageRatio(blocked, 100-w.strength)
	}

	if redux < minRedux {
		redux = minRedux
	}

	if w.durability-redux <= 0 {
		w.durability = 0
		w.dequip()
	} else {
		w.durability -= redux
	}

}

func (p *person) getWeaponBluntness() int {
	return p.getCurrentWeapon().bluntness
}

func (p *person) getWeaponSharpness() int {
	return p.getCurrentWeapon().sharpness
}

// armorEquip is an armor instance.
type armorEquip struct {
	equipment
	armor     armor
	strength  int
	hardness  int
	dampening int
}

func (a *armorEquip) getHardness() int {
	return calcDamageRatio(a.hardness, a.durability)
}

func (a *armorEquip) getDampening() int {
	return calcDamageRatio(a.dampening, a.durability)
}
