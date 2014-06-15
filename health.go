package characters

type Health struct {
	Head    int
	Chest   int
	Stomach int
	Back    int
	LArm    int
	RArm    int
	LHand   int
	RHand   int
	LLeg    int
	RLeg    int
	LFoot   int
	RFoot   int
}

func NewHealth() Health {
	return Health{
		Head:    100,
		Chest:   100,
		Stomach: 100,
		Back:    100,
		LArm:    100,
		RArm:    100,
		LHand:   100,
		RHand:   100,
		LLeg:    100,
		RLeg:    100,
		LFoot:   100,
		RFoot:   100,
	}
}

//
func (h *Health) MovementCapacity() float32 {
	general := h.Head + h.Chest + h.Stomach + h.Back + h.LLeg + h.RLeg + h.LFoot + h.RFoot
	capacity := float32(general) / 800

	// Check for broken bones
	if h.LLeg <= 50 {
		capacity /= 2
	}

	if h.RLeg <= 50 {
		capacity /= 2
	}

	if h.LFoot <= 50 {
		capacity /= 2
	}

	if h.RFoot <= 50 {
		capacity /= 2
	}

	return capacity

}
