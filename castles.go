package main

import "fmt"
import "strings"

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
	alchemy  castle_workshop = '⚗'      // alembic
	military castle_workshop = '⚔'      // crossed swords
	economy  castle_workshop = '⊚'      // coin
	industry castle_workshop = '⚒'      // mallet and pick
	intrigue castle_workshop = "\u2709" // a sealed letter
	arts     castle_workshop = '♪'      // some artsy doilie
)

type wall_thickness uint8

const (
	thick wall_thickness = iota
	double
	dashed
	thin
	broken
)

type castle struct {
	style     castle_style
	size      castle_size
	towers    []castle_workshop
	gates     []cardinal
	thickness wall_thickness
}

// castle.towerString builds a visual representation of a tower in a castle.
func (c castle) towerString() string {

	// Height of the tower
	var v_count, h_count int

	// Different sizes of castle have different sized towers
	h_count, v_count = c.towerDims()

	// Amount of non-edge, non-midpoint pieces in each row. Defined as the
	// tower's width minus two for edges minus one for the midpoint
	o_count := h_count - 2 - 1

	// "Pieces" (runes) of the outer horizontal and vertical edges of the tower
	var p_hout, p_vout rune

	// "Pieces" (runes) of the inner horizontal and vertical edges of the tower
	var p_hin, p_vin rune

	// The runes of the corner pieces of each tower: top-left, top-right,
	// bottom-left, bottom-right
	var corners [4]rune

	// The runes of the midpoints of each tower: top, right, bottom, left
	var midpoints [4]rune

	// Each style of tower has a different set of input runes
	switch c.style {
	case mexican:
		corners = [4]rune{'╭', '╮', '╰', '╯'}
		midpoints = [4]rune{'╥', '╡', '╨', '╞'}
		p_hout = '─'
		p_vout = '│'
		p_hin = '═'
		p_vin = '║'
	case japanese:
		corners = [4]rune{'┯', '┯', '┷', '┷'}
		midpoints = [4]rune{'━', '┤', '━', '├'}
		p_hout = '━'
		p_vout = '│'
		p_hin = ' '
		p_vin = ' '
	case american:
		corners = [4]rune{'╒', '╕', '╘', '╛'}
		midpoints = [4]rune{'═', '┃', '═', '┃'}
		p_hout = '═'
		p_vout = '┃'
		p_hin = ' '
		p_vin = ' '
	case w_euro:
		square := '\u25fb'
		corners = [4]rune{square, square, square, square}
		midpoints = [4]rune{square, '┇', square, '┇'}
		p_hout = '╍'
		p_vout = '┇'
		p_hin = ' '
		p_vin = ' '
	}

	// Building a string representation of the tower
	var t_line string
	var t_lines []string

	for row := v_count; row >= 1; row-- {
		var left_edge, right_edge, mid, others rune

		switch row {

		// first row
		case v_count:
			left_edge = corners[0]
			right_edge = corners[1]
			others = p_hout
			mid = midpoints[0]

		// last row
		case 1:
			left_edge = corners[2]
			right_edge = corners[3]
			others = p_hout
			mid = midpoints[2]

		// middle row
		case v_count/2 + 1:
			left_edge = midpoints[3]
			right_edge = midpoints[1]
			others = p_hin
			mid = ' ' //TODO: put things in towers

		default:
			left_edge = p_vout
			right_edge = p_vout
			others = ' '
			mid = p_vin

		}

		// Put the pieces together for each row
		t_line = string(left_edge)
		t_line += strings.Repeat(string(others), o_count/2)
		t_line += string(mid)
		t_line += strings.Repeat(string(others), o_count/2)
		t_line += string(right_edge)

		// Put the current row's combined pieces in a list of rows
		t_lines = append(t_lines, t_line)
	}

	return strings.Join(t_lines, "\n")
}

// castle.towerDims determines the dimensions of the towers of a castle based
// on its size, which isn't numeric
func (c castle) towerDims() (h_count, v_count int) {
	switch c.size {
	case small:
		v_count = 3
		h_count = 5
	case large:
		v_count = 5
		h_count = 9
	case enormous:
		v_count = 7
		h_count = 13
	}

	return
}

func (c castle) wallDims() (h_count, v_count int) {
	h_count, v_count = c.towerDims()

	h_count *= 3
	v_count *= 3

	return
}

func (c castle) dims() (h_count, v_count int) {
	h_tower, v_tower := c.towerDims()
	h_wall, v_wall := c.wallDims()

	h_count = h_tower*2 + h_wall
	v_count = v_tower*2 + v_wall

	return
}

// castle.String builds a string representation of the entire castle, towers,
// walls, etc. all included.
func (c castle) String() string {

	// Different sizes of castle have different length walls and towers
	h_count, v_count := c.wallDims()
	h_tower, v_tower := c.towerDims()

	// Determine gap between tower edge and wall
	h_spaceOutside := (h_tower - 5) / 4
	v_spaceOutside := (v_tower - 1) / 4

	// Determine space between walls
	h_spaceBetween := h_tower - h_spaceOutside*2 - 2
	//v_spaceBetween := v_tower - v_spaceOutside*2 - 2

	// Horizontal and vertical "pieces" (runes) of the wall
	var p_v, p_h rune

	// Each wall thickness has its own mini-style
	switch c.thickness {
	case thick:
		p_v = '┃'
		p_h = '━'
	case double:
		p_v = '║'
		p_h = '═'
	case dashed:
		p_v = '┇'
		p_h = '╍'
	case thin:
		p_v = '│'
		p_h = '─'
	case broken:
		p_v = '╵'
		p_h = '╴'
	}

	/*
		// But all gates look the same
		p_gate_v := rune('⸬')
		p_gate_h := rune('⸬')
	*/

	lines_tower := strings.Split(c.towerString(), "\n")

	// Copy lines_tower to a new slice `lines`, but add enough rows to fit the
	// whole castle
	_, castleLines := c.dims()
	lines := make([]string, castleLines)
	for i := 0; i < len(lines_tower); i++ {
		lines[i] = lines_tower[i]
	}

	// Begin to construct each row
	row := 0

	// Construct top towers and north walls
	first := row
	last := row + v_tower
	for ; row < last; row++ {

		// determine whether to add wall or space
		var s string

		if row-first == v_spaceOutside || last-row-1 == v_spaceOutside {
			// we're on top of the wall
			s = string(p_h)
		} else {
			// we're either outside or inside the wall
			s = " "
		}

		// Add wall or space
		lines[row] += strings.Repeat(s, h_count)

		// Add a line of right tower
		lines[row] += lines_tower[row-first]
	}

	// Add vertical walls on both sides, with space between
	last = row + v_count
	for ; row < last; row++ {

		// west wall: outsideSpace, outside, insideSpace, inside, outsideSpace
		lines[row] = strings.Repeat(" ", h_spaceOutside)
		lines[row] += string(p_v)
		lines[row] += strings.Repeat(" ", h_spaceBetween)
		lines[row] += string(p_v)
		lines[row] += strings.Repeat(" ", h_spaceOutside)

		// inside of castle
		lines[row] += strings.Repeat(" ", h_count)

		// east wall: outsideSpace, outside, space, inside, outsideSpace
		lines[row] += strings.Repeat(" ", h_spaceOutside)
		lines[row] += string(p_v)
		lines[row] += strings.Repeat(" ", h_spaceBetween)
		lines[row] += string(p_v)
		lines[row] += strings.Repeat(" ", h_spaceOutside)

	}

	// Construct bottom towers and south walls
	first = row
	last = row + v_tower
	for ; row < last; row++ {

		// Add a line of left tower
		lines[row] = lines_tower[row-first]

		// determine whether to add wall or space
		var s string

		if row-first == v_spaceOutside || last-row-1 == v_spaceOutside {
			// we're on top of the wall
			s = string(p_h)
		} else {
			// we're either outside or inside the wall
			s = " "
		}

		// Add wall or space
		lines[row] += strings.Repeat(s, h_count)

		// Add a line of right tower
		lines[row] += lines_tower[row-first]
	}

	return strings.Join(lines, "\n")

}

func main() {
	var c castle
	sizes := []castle_size{small, large, enormous}
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
				gates:     []cardinal{east},
				thickness: thin,
			}
			fmt.Println(c.String())
		}
	}
}
