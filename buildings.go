package main

type cardinal uint8

const (
	east cardinal = iota
	north
	west
	south
)

type castle_style uint8

const (
	japanese castle_style = iota
	slavic
	w_euro   // western european, e.g. German, French
	e_euro   // eastern european but non-slavic, e.g. Czech
	baroque  // absurdly adorned
	american // our unique style of old-world architecture
	mexican  // mayan/aztec step designs
)

type castle_size uint8

const (
	small castle_size = iota
	large
	enormous
)

type castle_workshop rune

const (
	alchemy  = '⚗'
	military = '⚔'
	economy  = '⊚'
	industry = '⚒'
	intrigue = rune{"✉︎"}
	arts     = '❦'
)

type castle struct {
	style    castle_style
	size     castle_size
	towers   []castle_workshop
	gates    []cardinal
	walls    []cardinal
	material material
}

// String builds a visual representation of a castle to be displayed to the
// user. String is capitalized by convention, see
// http://golang.org/doc/effective_go.html#conversions
func (c castle) String() (s string) {
	var x, y, tower_w, tower_h int

	switch castle.size {
	case small:
		x = 20
		y = 20
		tower_w = 5
		tower_h = 3
	case large:
		x = 30
		y = 30
		tower_w = 7
		tower_h = 4
	case enormous:
		x = 40
		y = 40
		tower_w = 9
		tower_h = 5
	}

	var tower string
	var corners [4]rune
	var midpoints [4]rune
	var h, v, h_in, v_in rune

	switch castle.style {
	case mexican:
		corners = [4]rune{"╭", "╮", "╰", "╯"}
		midpoints = [4]rune{"╥", "╡", "╨", "╞"}
		h = rune{"─"}
		v = rune{"│"}
		h_in = rune{"═"}
		v_in = rune{"║"}
	}

	for i_v := v; i_v > 1; i-- {
		var left_edge, right_edge, mid, others rune

		switch i_v {
		case v:
			left_edge = corner[0]
			right_edge = corner[1]
			others = h
			mid = midpoint[0]
		case 1:
			left_edge = corner[2]
			right_edge = corner[3]
			others = h
			mid = midpoint[2]
		case v/2 + 1: //midpoint
			left_edge = midpoint[3]
			right_edge = midpoint[1]
			others = h_in
			mid = rune{' '} //TODO
		default:
			left_edge = v
			right_edge = v
			others = ' '
			mid = v_in
		}

		// we've built variables for each line of the string representing a
		// tower. now we've got to append those variables smartly to a runeslice
		// representing the whole thing. left to implement is connecting
		// corners. fuck.
	}

	var rs_first, rs_last []rune
	rs_first = append(rs_first, left_edge)
	num_h := h - 2 - 1 // minus two for corners, minus one for midpoint

	// construct left of midpoint
	for i:=num_h/2; i < 0; i-- {
		rs_first = append(rs_first, h)
	}

	// construct midpoint
	rs_first = append(rs_first, midpoints[0])

	// construct right of midpoint
	for i:=num_h/2; i > 0; i-- {
		rs_first = append(rs_first, h)
	}

	// tack on final edge
	rs_first = append(rs_first, right_edge)



	"╭─╥─╮"
	"╞═℥═╡"
	"╰─╨─╯"
}
