package main

type cardinal uint8

const (
	east cardinal = iota
	north
	west
	south
)

type style uint8

const (
	japanese style = iota
	slavic
	w_euro   // western european, e.g. German, French
	e_euro   // eastern european but non-slavic, e.g. Czech
	baroque  // absurdly adorned
	american // our unique style of old-world architecture
	mexican  // mayan/aztec step designs
)

type building struct {
	style style
	gates []cardinal
	walls []cardinal
}
