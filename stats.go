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

func getFightExp(p profession) bool {
	switch {
	case p == soldier:
		return true
	case p == officer:
		return true
	}
	return false
}

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

func getToolSkill(p profession) int {
	return genSkill(getDist(getToolExp(p)))
}

func getFightSkill(p profession) int {

	return genSkill(getDist(getFightExp(p)))

}

func getPhysAtt(p profession) int {
	return genSkill(getDist(getPhysExp(p)))
}

func getIntAtt(p profession) int {
	return genSkill(getDist(getIntExp(p)))
}

type skills struct {
	weapons  map[weapon]int
	tools    map[tool]int
	armorUse map[armor]int
}

type attributes struct {
	strength     int
	agility      int
	accuracy     int
	intelligence int
	stamina      int
	social       int
}

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
