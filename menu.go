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
	midh := h/2 - 1
	lpad := (w - len(msg)) / 2
	tbprint(lpad, midh, fg, bg, msg)
}

func drawTinLogo(offset bool) {
	tbph := func(m string, h int) {
		tbcenterwidth(h, termbox.ColorWhite, termbox.ColorDefault, m)
	}
	_, h := termbox.Size()
	midh := h/2 - 1

	var top, bottom string
	if offset {
		top += " "
	} else {
		bottom += " "
	}

	tbph(top+"✰ ✰ ✰ ✰ ✰", midh-2)
	tbph(" T.I.N.! ", midh)
	tbph(bottom+"✰ ✰ ✰ ✰ ✰", midh+2)

	tbph("(ESC exits)", h-1)
}

func checkBoolChan(c chan bool) bool {
	b := <-c
	c <- b

	return b
}

func runGameLoop() {
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

	offset_chan := make(chan bool, 1)
	offset_chan <- false
	go func() {
		for {
			b := <-offset_chan
			offset_chan <- !b
			time.Sleep(1 * time.Second)
		}
	}()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawTinLogo(checkBoolChan(offset_chan))
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
				case termbox.KeySpace:
					// do nothing
				}
			}
		default:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			drawTinLogo(checkBoolChan(offset_chan))
			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
