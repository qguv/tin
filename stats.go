package main

import "math/rand"

// getDist returns skill probability distribution based on wether
// a character is experienced or not.
// Index corresponds to skill level, and value at index is the
// minimum roll needed to obtain that skill level.
// ex: dist[5] = 75 means a roll of 75 is need to have level 5
func getDist(skilled bool) [10]int {
	if skilled {
		return [10]int{5, 15, 25, 40, 60, 75, 90, 93, 95, 98}
	} else {
		return [10]int{40, 60, 70, 80, 85, 90, 94, 96, 98, 100}
	}
}

// genSkill takes in a probability distribution and rolls for
// skill value.
func genSkill(dist [10]int) int {
	roll := int(rand.Int31n(100))
	for result := 0; result < 10; result++ {
		if roll <= dist[result] {
			return result
		}
	}
	return 10
}

// getFightExp is an ugly function that keeps track of which
// professions have experience in fighting.
func getFightExp(p profession) bool {
	switch {
	case p == soldier:
		return true
	case p == officer:
		return true
	}
	return false
}

// getToolExp another ugly exp func.
func getToolExp(p profession) bool {
	switch {
	case p == laborer:
		return true
	case p == farmer:
		return true
	case p == craftsman:
		return true
	}
	return false
}

// getPhysExp another ugly exp func.
func getPhysExp(p profession) bool {
	switch {
	case p == laborer:
		return true
	case p == farmer:
		return true
	case p == soldier:
		return true
	case p == officer:
		return true
	}
	return false
}

// getIntExp another ugly exp function.
func getIntExp(p profession) bool {
	switch {
	case p == merchant:
		return true
	case p == officer:
		return true
	case p == scientist:
		return true
	case p == bureaucrat:
		return true
	case p == noble:
		return true
	}
	return false
}

// getToolSkill generates a starting tool skill for a profession.
func getToolSkill(p profession) int {
	return genSkill(getDist(getToolExp(p)))
}

// getFightSkill generates a starting fight (weapon/armor) skill
// for a profession.
func getFightSkill(p profession) int {

	return genSkill(getDist(getFightExp(p)))

}

// getPhysAtt generates a starting physical attribute for a profession.
func getPhysAtt(p profession) int {
	return genSkill(getDist(getPhysExp(p)))
}

// getIntAtt generates a starting intelligence attribute for a profession.
func getIntAtt(p profession) int {
	return genSkill(getDist(getIntExp(p)))
}

// skills represents a characters skill at using equipment.
// 0 - 10
type skills struct {
	weapons  map[weapon]int
	tools    map[tool]int
	armorUse map[armor]int
}

// attributes represent a characters personal traits.
// 0 - 10
type attributes struct {
	strength     int
	agility      int
	accuracy     int
	intelligence int
	stamina      int
	social       int
}

// newSkills returns initialized skills for a profession.
func newSkills(p profession) skills {
	weapons := make(map[weapon]int)
	for x := 0; x < numWeapons; x++ {
		weapons[weapon(x)] = getFightSkill(p)
	}
	tools := make(map[tool]int)
	for x := 0; x < numTools; x++ {
		tools[tool(x)] = getToolSkill(p)
	}
	armors := make(map[armor]int)
	for x := 0; x < numArmor; x++ {
		armors[armor(x)] = getFightSkill(p)
	}
	return skills{
		weapons:  weapons,
		armorUse: armors,
		tools:    tools,
	}
}

// newAttributes returns initialized attributes for a profession.
func newAttributes(p profession) attributes {
	return attributes{
		strength:     getPhysAtt(p),
		agility:      getPhysAtt(p),
		accuracy:     getPhysAtt(p),
		intelligence: getIntAtt(p),
		stamina:      getPhysAtt(p),
		social:       getIntAtt(p),
	}
}
