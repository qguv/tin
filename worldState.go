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

type manMadeType int

const (
	castleM = iota
	town
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

type manmadefeature interface {
	getOwner() *owner
	getTile() *tile
	getType() manMadeType
}

type manmade struct {
	name        string
	owner       *owner
	manMadeType manMadeType
	tile        *tile
}

type structureType int

const (
	admin = iota
	asylum
	barrack
	bath
	boardingHouse
	cemetary
	religious
	cistern
	coliseum
	corral
	fountain
	garden
	granery
	guildhouse
	hospital
	house
	infirmary
	inn
	library
	mill
	office
	plaza
	prison
	restaurant
	shop
	stable
	tavern
	tenemant
	theater
	university
	warehouse
	well
	workshop
)

type workshopType int

const (
	cobblers         = 87
	furnatureMakers  = 174
	furriers         = 240
	weavers          = 293
	basketMakers     = 335
	carpenders       = 377
	parchmentmakers  = 419
	potters          = 461
	wheelwrights     = 499
	jewlers          = 534
	masons           = 564
	bakers           = 594
	soapmakers       = 620
	chandlers        = 641
	coopers          = 661
	pastryMakers     = 680
	scabbardMakers   = 695
	silversmiths     = 710
	saddlers         = 723
	purseMakers      = 735
	blacksmiths      = 747
	goldsmiths       = 759
	toymakers        = 771
	artists          = 782
	leatherworkers   = 793
	ropemakers       = 803
	tanners          = 813
	buckleMakers     = 822
	cutlers          = 831
	fullers          = 840
	harnessMakers    = 849
	painters         = 858
	woodcarvers      = 866
	glassmakers      = 873
	instrumentmakers = 880
	locksmiths       = 887
	rugMakers        = 894
	sculptors        = 901
	shipmakers       = 913
	bookbinders      = 919
	fletchers        = 925
	brewers          = 931
	gloveMakers      = 937
	vintner          = 943
	skinners         = 953
	armorers         = 958
	weaponsmiths     = 963
	distillers       = 967
	illuminators     = 971
	perfumer         = 975
	tilers           = 979
	potionMakers     = 983
	clockMakers      = 986
	taxidermists     = 989
	vestmentmakers   = 992
	alchemists       = 994
	bellmakers       = 996
	dyeMakers        = 998
	inventors        = 1000
)

type shopType int

const (
	clothiers          = 97
	gorcers            = 194
	diarysellers       = 270
	launderers         = 346
	prostitutes        = 422
	furrierMerchants   = 498
	tailors            = 558
	barbers            = 607
	drapers            = 656
	flowersellers      = 705
	jewelers           = 745
	mercers            = 768
	engravers          = 790
	pawnbroker         = 812
	haberdasher        = 832
	wineMerchants      = 852
	tinkers            = 868
	butchers           = 883
	fishmongers        = 898
	woolMerchants      = 911
	beerMerchants      = 923
	herbalists         = 935
	spiceMerchants     = 947
	woodSellers        = 957
	brothelKeepers     = 965
	hayMerchants       = 973
	booksellers        = 979
	religiousSouvenirs = 985
	dentists           = 989
	navelOutfitters    = 993
	grainMerchants     = 996
	tobaccoMerchants   = 999
	magicMerchants     = 1000
)

//TODO officeType

type buildingStyle int

const (
	derelict = iota
	basic
	ornate
	imperial
)

type structure struct {
	structureType structureType
	buildingStyle buildingStyle
}

type districtType int

const (
	slum = iota
	gate
	sea
	industrial
	riverD
	militaryD
	craftsmen
	merchantD
	administrative
	market
	patriciate
)

type district struct {
	location
	districtType districtType
	structures   []structure
}

type settlement struct {
	manmade
	districts  []district
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
	location
	claimedBy      *owner
	biome          biome
	manmadefeature *manmadefeature
	roughness      int
	wildness       int
}
