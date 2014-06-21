package main

import "math"

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
	blood   int
	stamina int
}

// newPersonHealth initializes health to 100.
func newPersonHealth() health {
	parts := make([]bodyPartInstance, 12)
	for x := 0; x < numBodyParts; x++ {
		parts = append(parts, bodyPartInstance{
			bodyPart: bodyPart(x),
			health:   100,
		})
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
