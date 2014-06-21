package main

import "math/rand"

// equiped represents the equipment that a character is wearing / weilding.
type equipped struct {
	armor  map[armor]armorEquip
	weapon weaponEquip
	tool   toolEquip
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
	name       string
	age        int
	male       bool
	profession profession
	nobility   nobility
	father     *person
	mother     *person
	carried    []carriable
	holdings   []ownable
}

// getWeaponSkill returns the characters skill in using his equipped weapon.
// returns 0 if nothing equipped.
func (p *person) getWeaponSkill() int {
	weapon := p.equipped.weapon.weapon
	return p.skills.weapons[weapon]
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
	return person{
		profession: p,
		skills:     newSkills(p),
		attributes: newAttributes(p),
		health:     newPersonHealth(),
	}
}
