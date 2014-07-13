package main

type arbor uint8

const (
	oak arbor = iota
	maple
	cherry
	pine
	ash
)

type crop uint8

const (
	wheat crop = iota
	fruit
	vegetables
	tobacco
	coffee
	cotton
)

type mineralOre uint8

const (
	stoneOre mineralOre = iota
	copperOre
	tinOre
	ironOre
	silverOre
	goldOre
	adamantineOre
)

type wildlife uint8

const (
	bear wildlife = iota
	buffalo
	stag
	boar
	elephant
)

type domesticated uint8

const (
	hound domesticated = iota
	sheep
	goat
	chicken
	cow
	horse
	camel
)
