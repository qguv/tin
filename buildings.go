package main

import "fmt"

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
	w_euro                // western european, e.g. German, French
	american              // our unique style of old-world architecture
	mexican               // mayan/aztec step designs
)

type castle_size uint8

const (
	small castle_size = iota
	large
	enormous
)

type castle_workshop rune

const (
	alchemy  = '⚗'      // alembic
	military = '⚔'      // crossed swords
	economy  = '⊚'      // coin
	industry = '⚒'      // mallet and pick
	intrigue = "\u2709" // a sealed letter
	arts     = '❦'      // some artsy doilie
)

type wall_weight uint8

const (
	broken wall_weight = iota
	dashed
	thin
	thick
	double
)

type castle struct {
	style  castle_style
	size   castle_size
	towers []castle_workshop
	gates  []cardinal
	weight wall_weight
}

// String builds a visual representation of a castle to be displayed to the
// user. String is capitalized by convention, see
// http://golang.org/doc/effective_go.html#conversions
func (c castle) String() string {
	var corners, midpoints [4]rune
	var h, v, h_in, v_in rune
	var towerHeight, towerWidth int

	switch c.style {
	case mexican:
		corners = [4]rune{'╭', '╮', '╰', '╯'}
		midpoints = [4]rune{'╥', '╡', '╨', '╞'}
		h = '─'
		v = '│'
		h_in = '═'
		v_in = '║'
	case japanese:
		corners = [4]rune{'┯', '┯', '┷', '┷'}
		midpoints = [4]rune{'━', '┤', '━', '├'}
		h = '━'
		v = '│'
		h_in = ' '
		v_in = ' '
	case american:
		corners = [4]rune{'╒', '╕', '╘', '╛'}
		midpoints = [4]rune{'═', '┃', '═', '┃'}
		h = '═'
		v = '┃'
		h_in = ' '
		v_in = ' '
	case w_euro:
		square := '\u25fb'
		corners = [4]rune{square, square, square, square}
		midpoints = [4]rune{square, '┇', square, '┇'}
		h = '╍'
		v = '┇'
		h_in = ' '
		v_in = ' '
	}

	switch c.size {
	case small:
		towerHeight = 3
		towerWidth = 5
	case large:
		towerHeight = 5
		towerWidth = 9
	case enormous:
		towerHeight = 7
		towerWidth = 13
	}

	num_others := towerWidth - 2 - 1 // minus two for corners, minus one for midpoint

	var tower [][]rune
	var t_line []rune

	for i_v := towerHeight; i_v >= 1; i_v-- {
		var left_edge, right_edge, mid, others rune

		switch i_v {
		case towerHeight: // top
			left_edge = corners[0]
			right_edge = corners[1]
			others = h
			mid = midpoints[0]
		case 1: // bottom
			left_edge = corners[2]
			right_edge = corners[3]
			others = h
			mid = midpoints[2]
		case towerHeight/2 + 1: // middle
			left_edge = midpoints[3]
			right_edge = midpoints[1]
			others = h_in
			mid = ' ' //TODO
		default: // others
			left_edge = v
			right_edge = v
			others = ' '
			mid = v_in
		}

		t_line = []rune{left_edge}

		for i := num_others / 2; i > 0; i-- {
			t_line = append(t_line, others)
		}

		t_line = append(t_line, mid)

		for i := num_others / 2; i > 0; i-- {
			t_line = append(t_line, others)
		}

		t_line = append(t_line, right_edge)

		tower = append(tower, t_line)
	}

	var s string
	for _, t := range tower {
		s += string(t) + "\n"
	}
	return s
}

func main() {
	var c castle
	sizes := []castle_size{small, large}
	styles := []castle_style{w_euro, japanese, mexican, american}
	for _, style := range styles {
		for _, size := range sizes {
			c = castle{
				style: style,
				size:  size,
				towers: []castle_workshop{
					alchemy,
					economy,
					military,
					arts,
				},
				gates:  []cardinal{east},
				weight: thin,
			}
			fmt.Printf(c.String())
		}
	}
}
