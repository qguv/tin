package main

type health struct {
	head    int
	chest   int
	stomach int
	back    int
	lArm    int
	rArm    int
	lHand   int
	rHand   int
	lLeg    int
	rLeg    int
	lFoot   int
	rFoot   int
}

func newHealth() health {
	return health{
		head:    100,
		chest:   100,
		stomach: 100,
		back:    100,
		lArm:    100,
		rArm:    100,
		lHand:   100,
		rHand:   100,
		lLeg:    100,
		rLeg:    100,
		lFoot:   100,
		rFoot:   100,
	}
}

//
func (h *health) movementCapacity() float32 {
	general := h.head + h.chest + h.stomach + h.back + h.lLeg + h.rLeg + h.lFoot + h.rFoot
	capacity := float32(general) / 800

	// Check for broken bones
	if h.lLeg <= 50 {
		capacity /= 2
	}

	if h.rLeg <= 50 {
		capacity /= 2
	}

	if h.lFoot <= 50 {
		capacity /= 2
	}

	if h.rFoot <= 50 {
		capacity /= 2
	}

	return capacity

}
