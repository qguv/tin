package main

import (
	"math"
	"strconv"
)

// bodyPart is an enum of general types of body parts.
type bodyPart int

const (
	head = iota
	chest
	rArm
	lArm
	rHand
	lHand
	rLeg
	lLeg
	rFoot
	lFoot
	numBodyParts
)

func (b bodyPart) String() string {
	switch b {
	case head:
		return "head"
	case chest:
		return "chest"
	case rArm:
		return "right-arm"
	case lArm:
		return "left-arm"
	case rHand:
		return "right-hand"
	case lHand:
		return "left-hand"
	case rLeg:
		return "right-leg"
	case lLeg:
		return "left-leg"
	case rFoot:
		return "right-foot"
	case lFoot:
		return "left-foot"
	default:
		return ""
	}
}

// bodyPartInstance is an instance of a body part with status.
type bodyPartInstance struct {
	bodyPart bodyPart
	armor    *armorEquip
	// out of 100
	health   int
	broken   bool
	detached bool
	infected bool
	severed  bool
}

// getBodyPart returns a corresponding bodyPartInstance
// given a bodyPart
func (h *health) getBodyPart(bp bodyPart) *bodyPartInstance {
	for index := 0; index < len(h.bodyParts); index++ {
		if h.bodyParts[index].bodyPart == bp {
			return &h.bodyParts[index]
		}
	}
	return nil
}

// health represents a characters health status.
type health struct {
	bodyParts []bodyPartInstance
	// out of 100
	blood       int
	stamina     int
	unconscious bool
}

func (h *health) checkHealth() {

	hd := h.getBodyPart(head)

	if hd.broken || hd.detached || hd.severed || hd.health == 0 {
		h.unconscious = true
	}

	hd = h.getBodyPart(chest)

	if hd.detached || hd.health == 0 {
		h.unconscious = true
	}

	if h.totalHealth() <= 50 {
		h.unconscious = true
	}

	if h.blood <= 50 {
		h.unconscious = true
	}

}

func (h health) String() string {
	var ans string
	for _, x := range h.bodyParts {
		ans += x.bodyPart.String() + ": " + strconv.Itoa(x.health) + "\n"
	}
	return ans
}

func (h *health) totalHealth() int {
	var general, disabled int
	for _, x := range h.bodyParts {
		general += x.health

		// extra weight to head and chest
		if x.bodyPart == head {
			general += 3 * x.health
		} else if x.bodyPart == chest {
			general += 2 * x.health
		}

		if x.broken || x.infected || x.severed || x.detached {
			disabled++
		}
	}
	capacity := float32(general) / (100*float32(len(h.bodyParts)) + 500)
	capacity /= float32(math.Exp2(float64(disabled)))

	return int(capacity*100 + .5)

}

// newPersonHealth initializes health to 100.
func newPersonHealth() health {
	parts := make([]bodyPartInstance, numBodyParts)
	for x := 0; x < numBodyParts; x++ {
		parts[x] = bodyPartInstance{
			bodyPart: bodyPart(x),
			health:   100,
		}
	}
	return health{
		bodyParts: parts,
		blood:     100,
		stamina:   100,
	}
}

var (
	movementCrucial = []bodyPart{
		head,
		rLeg,
		lLeg,
		lFoot,
		rFoot,
		chest,
	}
)

// movementCapacity calculates a characters movement penalty
// from taking damage. It first calculates overall health, and
// then applies further redutions for broken bones.
// It returns movement ability in decimal form. Ex: .5 = 50%
func (h *health) movementCapacity() float32 {
	general := 0
	disabled := 0
	for _, x := range h.bodyParts {
		general += x.health
		if x.broken || x.infected || x.severed || x.detached {
			disabled++
		}
	}

	capacity := float32(general) / (100 * float32(len(h.bodyParts)))

	capacity /= float32(math.Exp2(float64(disabled)))

	return capacity

}
