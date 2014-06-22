package main

import "testing"
import termbox "github.com/nsf/termbox-go"

func TestStar(t *testing.T) {
	termbox.Init()
	defer termbox.Close()

	s := newStarAt(0)
	for i := 0; i <= STAR_GENERATION_COUNT; i++ {
		s.advance()
	}
	if !s.dead {
		t.Fatal("star did not die after aging it too long")
	}
}
