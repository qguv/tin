package main

import "math/rand"

type equiped struct {
	armor  map[armor]armorEquip
	weapon weaponEquip
	tool   toolEquip
}

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

type person struct {
	skills
	attributes
	equiped
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

func (p *person) getWeaponSkill() int {
	weapon := p.equiped.weapon.weapon
	return p.skills.weapons[weapon]
}

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

func newPerson(p profession) person {
	return person{
		profession: p,
		skills:     newSkills(p),
		attributes: newAttributes(p),
		health:     newHealth(),
	}
}
