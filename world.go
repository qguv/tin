package main

import "math/rand"
import "time"

const (
	WORLDWIDTH  = 100
	WORLDHEIGHT = 100
)

// totally subject to change
// just silly values for now
type world [WORLDHEIGHT][WORLDWIDTH]simpleTile
type worldview [][]simpleTile

func testerWorld() (theWorld world) {
	var b biome
	var t simpleTile
	rand.Seed(time.Now().UTC().UnixNano())

	for y := 0; y < WORLDHEIGHT; y++ {
		for x := 0; x < WORLDWIDTH; x++ {
			b = biome(rand.Intn(int(biome_count)))
			t = simpleTile{b}
			theWorld[y][x] = t
			mimicCell(x, y, t.Cell())
		}
	}

	return
}
