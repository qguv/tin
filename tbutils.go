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

// inKeyGroup tests whether a termbox event's reported key or rune is in a
// slice of matching keys or runes.
func inKeyGroup(k termbox.Key, r rune, keys []termbox.Key, runes []rune) bool {
	if r == rune(0) {
		return contains(keys, k)
	} else {
		return contains(runes, r)
	}
}
