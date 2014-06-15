package characters

import "math/rand"

type Equiped struct {
	Armor  map[Armor]ArmorEquip
	Weapon WeaponEquip
	Tool   ToolEquip
}

type Profession int

const (
	Vagrant = iota
	Laborer
	Farmer
	Soldier
	Craftsman
	Merchant
	Officer
	Scientist
	Bureaucrat
	Noble
	numProfessions
)

type Nobility int

const (
	Commoner = iota
	Knight
	Baron
	Viscount
	Count
	Duke
	Archduke
	Prince
	King
)

type Person struct {
	Skills
	Attributes
	Equiped
	Health
	Name       string
	Age        int
	Male       bool
	Profession Profession
	Nobility   Nobility
	Father     *Person
	Mother     *Person
	Carried    []Carriable
	Holdings   []Ownable
}

func (p *Person) GetWeaponSkill() int {
	weapon := p.Equiped.Weapon.Weapon
	return p.Skills.Weapons[weapon]
}

func RandPerson() Person {
	prof := Profession(rand.Int31n(numProfessions))
	person := NewPerson(prof)

	if rand.Int31n(1) == 1 {
		person.Male = true
	} else {
		person.Male = false
	}

	person.Age = int(rand.Int31n(22) + 18)

	return person
}

func NewPerson(p Profession) Person {
	return Person{
		Profession: p,
		Skills:     NewSkills(p),
		Attributes: NewAttributes(p),
		Health:     NewHealth(),
	}
}
