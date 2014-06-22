package main

import termbox "github.com/nsf/termbox-go"
import "time"
import "unicode/utf8"
import "math/rand"

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func tbcenterwidth(y int, fg, bg termbox.Attribute, msg string) {
	w, _ := termbox.Size()
	lpad := (w - utf8.RuneCountInString(msg)) / 2
	tbprint(lpad, y, fg, bg, msg)
}

func drawTinLogo() {
	tbph := func(m string, h int) {
		tbcenterwidth(h, termbox.ColorWhite, termbox.ColorDefault, m)
	}
	_, h := termbox.Size()
	midh := h / 2

	logo_raw := `
                       
 ╔═══════════════════╗ 
  ║ ⚓︎   ☠   ♘   ⚒   ⚇ ║ 
 ║    ╭─╮            ║ 
 ║ ♖ ╭╯ ╰┬─┬─────╮ ⎈ ║ 
 ║   ╰╮ ╭┤ │ ┌─╮ │   ║ 
 ║ ♪  │ ││ │ │ │ │ ♔ ║ 
 ║    └─┘└─┴─┘ └─┘   ║ 
  ║ ⚗   ◎   ⚔   ✙   ✉︎ ║ 
 ╚═══════════════════╝ 
                       
`
	logo := stringToLines(logo_raw)
	logo_h := len(logo)

	top := midh - logo_h/2

	for i, line := range logo {
		tbph(line, top+i)
	}

	adjustMessage_raw := `
                            
 adjust text size until 
 all characters are legible 
                            
	`
	adjustMessage := stringToLines(adjustMessage_raw)
	for i, s := range adjustMessage {
		tbph(s, i+1)
	}

	tbph("ESC exits", h-1)
}

type star struct {
	x, y       int
	generation int
	dead       bool
}

const STAR_GENERATION_COUNT int = 9

func newStarAt(generation int) star {
	h, w := termbox.Size()
	y := rand.Intn(w)
	x := rand.Intn(h)

	return star{
		x:          x,
		y:          y,
		generation: generation,
	}
}

func newStar() star {
	g := rand.Intn(STAR_GENERATION_COUNT)
	return newStarAt(g)
}

func (s *star) advance() {
	s.generation++
	if s.generation >= STAR_GENERATION_COUNT {
		s.dead = true
	}
}

func (s star) show() {
	thinEightPointStar := '\u2734'
	glyphs := []rune{
		'✢', '✧',
		'✶', '✵', '❃',
		'✷', thinEightPointStar,
		'✧', '◦',
	}

	r := glyphs[s.generation]
	termbox.SetCell(s.x, s.y, r, termbox.ColorWhite, termbox.ColorDefault)
}

func showStars(stars []star) {
	for _, star := range stars {
		star.show()
	}
}

func makeStars(count int) []star {
	stars := make([]star, count)
	for i := 0; i < count; i++ {
		stars[i] = newStar()
	}
	return stars
}

func advanceStars(stars []star) {
	for {
		this := rand.Intn(len(stars))
		s := &stars[this]
		s.advance()
		if s.dead {
			stars[this] = newStarAt(0)
		}

		time.Sleep(3 * time.Millisecond)
	}
}

func runGameLoop() {
	rand.Seed(time.Now().UTC().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	stars := makeStars(20)
	go advanceStars(stars)

gameLoop:
	for {
		select {
		case ev := <-event_queue:
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					break gameLoop
				case termbox.KeyCtrlC:
					break gameLoop
				case termbox.KeySpace:
					// do nothing
				}
			}
		default:
			showStars(stars)
			drawTinLogo()
			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
