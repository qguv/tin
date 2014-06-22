package main

import "math/rand"

// equiped represents the equipment that a character is wearing / weilding.
type equipped struct {
	armor  map[armor]*armorEquip
	weapon *weaponEquip
	tool   *toolEquip
}

func (p *person) carryCapacity() int {
	return calcDamageRatio(p.weight, 50+p.attributes.strength)
}

func (p *person) carry(i carriable) {
	if p.carryCapacity()-p.carriedWeight <= i.getWeight() {
		p.carried = append(p.carried, i)
		p.carriedWeight += i.getWeight()
	}
}

func (p *person) equipArmor(a *armorEquip) {
	p.equipped.armor[a.armor].dequip()
	p.equipped.armor[a.armor] = a

	for x, _ := range p.bodyParts {
		if p.bodyParts[x].bodyPart.armorType() == a.armor {
			p.bodyParts[x].armor = a
		}
	}
}

func (p *person) equipWeapon(w *weaponEquip) {
	p.equipped.tool = nil
	p.equipped.weapon = w
}

func (p *person) equipTool(t *toolEquip) {
	p.equipped.weapon = nil
	p.equipped.tool = t
}

// profession represents a characters general catagory of work.
type profession int

const (
	vagrant = iota
	laborer
	farmer
	soldier
	craftsman
	merchant
	officer
	scientist
	bureaucrat
	noble
	numProfessions
)

// nobility is the noble title held by a character.
type nobility int

const (
	commoner = iota
	knight
	baron
	viscount
	count
	duke
	archduke
	prince
	king
)

// person represents a human type character.
type person struct {
	skills
	attributes
	equipped
	health
	name          string
	age           int
	male          bool
	weight        int
	height        int
	carriedWeight int
	profession    profession
	nobility      nobility
	father        *person
	mother        *person
	carried       []carriable
	holdings      []ownable
}

// getWeaponSkill returns the characters skill in using his equipped weapon.
// returns 0 if nothing equipped.
func (p *person) getWeaponSkill() int {
	return p.skills.weapons[p.getCurrentWeapon().weapon]
}

func (p *person) getCurrentWeapon() *weaponEquip {
	if p.equipped.weapon == nil {
		return &weaponEquip{
			equipment: equipment{
				owner: p,
			},
			maxDamage: 20,
			bluntness: 100,
		}
	} else {
		return p.equipped.weapon
	}
}

// randPerson randomly chooses a profession,
// and then generates a character of that profession.
func randPerson() person {
	prof := profession(rand.Int31n(numProfessions))
	person := newPerson(prof)

	if rand.Int31n(1) == 1 {
		person.male = true
	} else {
		person.male = false
	}

	person.age = int(rand.Int31n(22) + 18)

	return person
}

// newPerson initializes a person with skill distribution
// based on chosen profession.
func newPerson(p profession) person {
	person := person{
		profession: p,
		skills:     newSkills(p),
		attributes: newAttributes(p),
		health:     newPersonHealth(),
	}

	return person
}
