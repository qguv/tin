package main

import termbox "github.com/nsf/termbox-go"
import "time"
import "unicode/utf8"
import "strings"

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
 ⚓︎   ☠   ♘   ⚒   ⚇
   ╭─╮           
♖ ╭╯ ╰┬─┬─────╮ ⎈
  ╰╮ ╭┤ │ ┌─╮ │  
♪  │ ││ │ │ │ │ ♔
   └─┘└─┴─┘ └─┘  
 ⚗   ◎   ⚔   ✙   ✉︎
`
	logo := strings.Split(logo_raw, "\n")
	logo_h := len(logo)
	logo = logo[1:logo_h]
	logo_h = len(logo)

	top := midh - logo_h/2

	for i, line := range logo {
		tbph(line, top+i)
	}

	tbph("adjust text size until", 1)
	tbph("all characters are legible", 2)
	tbph("ESC exits", h-1)
}

func runGameLoop() {
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawTinLogo()
	termbox.Flush()

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
			case termbox.EventResize:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				drawTinLogo()
				termbox.Flush()
			}
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
