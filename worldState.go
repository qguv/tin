package main

type biome int

const (
	desert = iota
	mountians
	plains
	lowlands
	tundra
	forest
	jungle
	hills
	swamp
	// water related
	coast
	ocean
	river
	lake
	haven
)

type constructionType int

const (
	// cant use castle, already declared
	// castle = iota
	town = iota
	village
	keep
	port
	farm
	mine
	camp
)

type location struct {
	xcord int
	ycord int
}

type owner interface {
	getHoldings() []ownable
}

type construction interface {
	getOwner() *owner
	getTile() *tile
	getType() constructionType
}

type construct struct {
	name            string
	owner           *owner
	construtionType constructionType
	tile            *tile
}

type buildings struct {
	lumbermill int
	woodshop   int
	masonshop  int
	forge      int
	smelter    int
}

type settlement struct {
	construct
	buildings
	population []popGroup
	treasury   int
}

// a group of people, the person represents the average person
// in the group
type popGroup struct {
	person
	population int
}

type naturalResources struct {
	trees      int
	stone      int
	tin        int
	copper     int
	iron       int
	adamantine int
}

type tile struct {
	naturalResources
	claimedBy    *owner
	location     location
	biome        biome
	construction *construction
	roughness    int
	wildness     int
}
