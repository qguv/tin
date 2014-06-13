package main

import "github.com/nsf/termbox-go"
import "time"
import "unicode/utf8"

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

func tbcenter(fg, bg termbox.Attribute, msg string) {
	w, h := termbox.Size()
	midh := h / 2 - 1
	lpad := (w - len(msg)) / 2
	tbprint(lpad, midh, fg, bg, msg)
}

func drawTinLogo() {
	tbph := func(m string, h int) {
		tbcenterwidth(h, termbox.ColorWhite, termbox.ColorDefault, m)
	}
	_, h := termbox.Size()
	midh := h / 2 - 1

	tbph("✰ ✰ ✰ ✰ ✰", midh - 2)
	tbph(" T.I.N.! ", midh)
	tbph("✰ ✰ ✰ ✰ ✰", midh + 2)

	tbph("(ESC exits)", h - 1)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

gameLoop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					break gameLoop
				case termbox.KeySpace:
					// do nothing
				}
			}
		default:
			drawTinLogo()
			time.Sleep(10 * time.Millisecond)
		}
		termbox.Flush()
	}
}
