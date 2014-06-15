package main

import "math/rand"

//
func getDist(skilled bool) [10]int {
	if skilled {
		return [10]int{5, 15, 25, 40, 60, 75, 90, 93, 95, 98}
	} else {
		return [10]int{40, 60, 70, 80, 85, 90, 94, 96, 98, 100}
	}
}

func genSkill(dist [10]int) int {
	roll := int(rand.Int31n(100))
	for result := 0; result < 10; result++ {
		if roll <= dist[result] {
			return result
		}
	}
	return 10
}

func getFightExp(p Profession) bool {
	switch {
	case p == Soldier:
		return true
	case p == Officer:
		return true
	}
	return false
}

func getToolExp(p Profession) bool {
	switch {
	case p == Laborer:
		return true
	case p == Farmer:
		return true
	case p == Craftsman:
		return true
	}
	return false
}

func getPhysExp(p Profession) bool {
	switch {
	case p == Laborer:
		return true
	case p == Farmer:
		return true
	case p == Soldier:
		return true
	case p == Officer:
		return true
	}
	return false
}

func getIntExp(p Profession) bool {
	switch {
	case p == Merchant:
		return true
	case p == Officer:
		return true
	case p == Scientist:
		return true
	case p == Bureaucrat:
		return true
	case p == Noble:
		return true
	}
	return false
}

func getToolSkill(p Profession) int {
	return genSkill(getDist(getToolExp(p)))
}

func getFightSkill(p Profession) int {

	return genSkill(getDist(getFightExp(p)))

}

func getPhysAtt(p Profession) int {
	return genSkill(getDist(getPhysExp(p)))
}

func getIntAtt(p Profession) int {
	return genSkill(getDist(getIntExp(p)))
}

type Skills struct {
	Weapons  map[Weapon]int
	Tools    map[Tool]int
	ArmorUse map[Armor]int
}

type Attributes struct {
	Strength     int
	Agility      int
	Accuracy     int
	Intelligence int
	Stamina      int
	Social       int
}

func NewSkills(p Profession) Skills {
	weapons := make(map[Weapon]int)
	for x := 0; x < numWeapons; x++ {
		weapons[Weapon(x)] = getFightSkill(p)
	}
	tools := make(map[Tool]int)
	for x := 0; x < numTools; x++ {
		tools[Tool(x)] = getToolSkill(p)
	}
	armor := make(map[Armor]int)
	for x := 0; x < numArmor; x++ {
		armor[Armor(x)] = getFightSkill(p)
	}
	return Skills{
		Weapons:  weapons,
		ArmorUse: armor,
		Tools:    tools,
	}
}

func NewAttributes(p Profession) Attributes {
	return Attributes{
		Strength:     getPhysAtt(p),
		Agility:      getPhysAtt(p),
		Accuracy:     getPhysAtt(p),
		Intelligence: getIntAtt(p),
		Stamina:      getPhysAtt(p),
		Social:       getIntAtt(p),
	}
}
