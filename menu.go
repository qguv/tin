package main

import termbox "github.com/nsf/termbox-go"
import "time"
import "unicode/utf8"
import "math/rand"

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

func drawCalibrateReminder() {
	_, h := termbox.Size()
	putUserMessage("SPACE to calibrate", h-2)
	putUserMessage("Ctrl-C to quit", h-1)
}

// drawTinLogo centers a logo on the screen for calibration.
func drawTinLogo() {
	_, h := termbox.Size()
	midh := h / 2

	adjustMessage_raw := `
                            
 adjust text size until 
 all characters are legible 
                            
	`
	adjustMessage := stringToLines(adjustMessage_raw)
	for i, s := range adjustMessage {
		putUserMessage(s, i+1)
	}

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

	top := midh - logo_h/2 + 1

	for i, line := range logo {
		putUserMessage(line, top+i)
	}

	putUserMessage("ESC exits", h-1)
}

type star struct {
	x, y       int
	generation int
	dead       bool
}

const STAR_GENERATION_COUNT int = 9

// newStarAt spits out a new star instance at a certain generation.
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

// newStar spits out a new star instance at a random generation.
func newStar() star {
	g := rand.Intn(STAR_GENERATION_COUNT)
	return newStarAt(g)
}

// *star.advance increases the generation of a certain star. If the *star
// becomes older than its maximum possible age as defined by
// STAR_GENERATION_COUNT, advance updates the *star indicating its death.
func (s *star) advance() {
	s.generation++
	if s.generation >= STAR_GENERATION_COUNT {
		s.dead = true
	}
}

// star.show displays the star on the screen at its current generation with the
// proper rune representing its age.
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

// showStars displays all stars in the slice at their current generation.
func showStars(stars []star) {
	for _, star := range stars {
		star.show()
	}
}

// makeStars generates a slice of stars of the given count. Since slices act as
// pointers, we only need to pass pointers around to methods that update state,
// so we're not returning anything but values here.
func makeStars(count int) []star {
	stars := make([]star, count)
	for i := 0; i < count; i++ {
		stars[i] = newStar()
	}
	return stars
}

// advanceStars repeatedly advances the value of a random star in the given
// slice. If the star dies, it is replaced by a new baby star at that position
// in the slice. Since this repeats ad infinatum, call this concurrently [go
// advanceStars(stars)].
func advanceStars(stars []star) {
	for {
		this := rand.Intn(len(stars))
		s := &stars[this]
		s.advance()
		if s.dead {
			stars[this] = newStarAt(0)
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func inKeyGroup(k termbox.Key, r rune, keys []termbox.Key, runes []rune) bool {
	if r == rune(0) {
		return contains(keys, k)
	} else {
		return contains(runes, r)
	}
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

// displayMainMenu does just that. Call this from main() if you want to run
// this part of the code.
func displayMainMenu() {
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

	stars := makeStars(6)
	go advanceStars(stars)

	var calibrate bool
	var selected int

gameLoop:
	for {
		select {
		case ev := <-event_queue:
			switch ev.Type {
			case termbox.EventKey:
				switch {
				case ev.Key == termbox.KeyCtrlC:
					break gameLoop
				case ev.Key == termbox.KeySpace:
					calibrate = !calibrate
				case ev.Key == termbox.KeyEsc:
					calibrate = false
				case isUpKey(ev.Key, ev.Ch):
					selected--
				case isDownKey(ev.Key, ev.Ch):
					selected++
				}
			}
		default:
			showStars(stars)
			if calibrate {
				drawTinLogo()
			} else {
				drawCalibrateReminder()
			}
			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
