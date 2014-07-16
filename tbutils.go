package main

import termbox "github.com/nsf/termbox-go"
import "unicode/utf8"

// putString displays a message horizontally, anchored on its left side by a
// given coördinate.
func putString(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

// putStringMidScreen displays a message horizontally, centered on the screen
func putStringMidScreen(fg, bg termbox.Attribute, msg string) {
	w, h := termbox.Size()
	y := h / 2
	lpad := (w - utf8.RuneCountInString(msg)) / 2
	putString(lpad, y, fg, bg, msg)
}

// putStringCentered displays a message horizontally, centered by width, at the
// given height (y).
func putStringCentered(y int, fg, bg termbox.Attribute, msg string) {
	w, _ := termbox.Size()
	lpad := (w - utf8.RuneCountInString(msg)) / 2
	putString(lpad, y, fg, bg, msg)
}

// putUserMessage displays a centered message to the user in white at the given
// height.
func putUserMessage(m string, h int) {
	m = smallText(m)
	putStringCentered(h, termbox.ColorWhite, termbox.ColorDefault, m)
}

// putMidScreen displays a completely centered message to the user in white.
func putMidScreen(m string) {
	putStringMidScreen(termbox.ColorDefault, termbox.ColorDefault, m)
}

// inKeyGroup tests whether a termbox event's reported key or rune is in a
// slice of matching keys or runes.
func inKeyGroup(k termbox.Key, r rune, keys []termbox.Key, runes []rune) bool {
	if r == rune(0) {
		return contains(keys, k)
	} else {
		return contains(runes, r)
	}
}

func isLeftKey(k termbox.Key, r rune) bool {
	leftKeys := []termbox.Key{
		termbox.KeyArrowLeft,
	}
	leftRunes := []rune{
		'h', 'H',
		'a', 'A',
	}
	return inKeyGroup(k, r, leftKeys, leftRunes)
}

func isRightKey(k termbox.Key, r rune) bool {
	rightKeys := []termbox.Key{
		termbox.KeyArrowRight,
	}
	rightRunes := []rune{
		'd', 'D',
		'l', 'L',
	}
	return inKeyGroup(k, r, rightKeys, rightRunes)
}

func isDownKey(k termbox.Key, r rune) bool {
	downKeys := []termbox.Key{
		termbox.KeyArrowDown,
		termbox.KeyPgdn,
	}
	downRunes := []rune{
		's', 'S',
		'j', 'J',
	}
	return inKeyGroup(k, r, downKeys, downRunes)
}

func isUpKey(k termbox.Key, r rune) bool {
	upKeys := []termbox.Key{
		termbox.KeyArrowUp,
		termbox.KeyPgup,
	}
	upRunes := []rune{
		'w', 'W',
		'k', 'K',
	}
	return inKeyGroup(k, r, upKeys, upRunes)
}

func tbClear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

const (
	Black      = termbox.ColorBlack
	DarkRed    = termbox.ColorRed
	DarkGreen  = termbox.ColorGreen
	DarkOrange = termbox.ColorYellow
	DarkBlue   = termbox.ColorBlue
	DarkBrown  = termbox.ColorMagenta
	DarkCyan   = termbox.ColorCyan
	Grey       = termbox.ColorWhite

	DarkGrey   = termbox.AttrBold | termbox.ColorBlack
	Red        = termbox.AttrBold | termbox.ColorRed
	Green      = termbox.AttrBold | termbox.ColorGreen
	DarkYellow = termbox.AttrBold | termbox.ColorYellow
	Blue       = termbox.AttrBold | termbox.ColorBlue
	Brown      = termbox.AttrBold | termbox.ColorMagenta
	Cyan       = termbox.AttrBold | termbox.ColorCyan
	White      = termbox.AttrBold | termbox.ColorWhite
)

func mimicCell(x, y int, c termbox.Cell) {
	termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
}
