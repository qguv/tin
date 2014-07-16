// World Explorer (a map)
// assumes a reference terminal of
// 84 (width) by 26 (height)

package main

import termbox "github.com/nsf/termbox-go"

func drawOutline() {
	x_count := 84
	y_count := 26
	w, h := termbox.Size()
	x_pad := (w - x_count) / 2
	y_pad := (h - y_count) / 2

	for y := y_pad; y < y_pad+26; y++ {
		for x := x_pad; x < x_pad+84; x++ {
			switch y {

			// top line
			case y_pad:
				termbox.SetCell(x, y, '-', termbox.ColorWhite, termbox.ColorDefault)

			// bottom line
			case y_pad + y_count - 1:
				termbox.SetCell(x, y, '-', termbox.ColorWhite, termbox.ColorDefault)

			default:
				switch x {

				// leftmost line
				case x_pad:
					termbox.SetCell(x, y, '|', termbox.ColorWhite, termbox.ColorDefault)

				// rightmost line
				case x_pad + x_count - 1:
					termbox.SetCell(x, y, '|', termbox.ColorWhite, termbox.ColorDefault)
				}
			}
		}
	}
}

func displayWorldExplorer() {
	termbox.Init()
	defer termbox.Close()

	tbClear()
	drawOutline()
	termbox.Flush()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	redraw := make(chan string)
	go func() {
		for {
			m := <-redraw
			tbClear()
			drawOutline()
			putMidScreen(m)
			termbox.Flush()
		}
	}()

termboxLoop:
	for {
		select {
		case ev := <-event_queue:
			switch ev.Type {
			case termbox.EventResize:
				redraw <- "resized"

			case termbox.EventKey:
				switch {

				case ev.Key == termbox.KeyCtrlC:
					break termboxLoop
					// TODO: die

				case ev.Key == termbox.KeyCtrlL:
					termbox.Sync()
					redraw <- "screen refreshed"

				case ev.Key == termbox.KeyEsc:
					break termboxLoop

				case isUpKey(ev.Key, ev.Ch):
					// TODO: move up
					redraw <- "up"

				case isDownKey(ev.Key, ev.Ch):
					// TODO: move down
					redraw <- "down"

				case isLeftKey(ev.Key, ev.Ch):
					// TODO: move left
					redraw <- "left"

				case isRightKey(ev.Key, ev.Ch):
					// TODO: move right
					redraw <- "right"

				default:
					redraw <- "unbound key"
				}

			default:
				redraw <- "something else"
			}
		}
	}
}
