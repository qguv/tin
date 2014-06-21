package main

import "testing"
import "fmt"

func TestCastleString(t *testing.T) {
	var c castle
	sizes := []castle_size{small, large, enormous}
	styles := []castle_style{wEuro_style, japanese_style, mexican_style, american_style}
	for _, style := range styles {
		for _, size := range sizes {
			c = castle{
				style: style,
				size:  size,
				towers: []castle_workshop{
					alchemy,
					economy,
					military,
					arts,
				},
				gates:     []cardinal{east},
				thickness: thin,
			}
			fmt.Println(c.String())
		}
	}
}
