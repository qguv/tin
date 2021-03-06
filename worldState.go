package main

import termbox "github.com/nsf/termbox-go"

type biome uint8

const (
	// low elevation (Norfolk, Holland, the tropics)
	desert biome = iota
	lowlands
	jungle

	// water-related (sea level)
	coast
	ocean
	river
	lake
	haven

	// mid elevation (Common)
	midlands
	tundra
	forest
	hills
	cliffs

	// high elevation (Charlottesville, West Virginia)
	highlands
	mountains

	// extreme (Colorado, Tibet) elevation
	fourteeners

	biome_count
)

var biomeCells = map[biome]termbox.Cell{
	desert:      termbox.Cell{'∾', White, DarkYellow},
	lowlands:    termbox.Cell{'⨪', Brown, DarkGreen},
	jungle:      termbox.Cell{'ꙟ', Cyan, DarkBrown},
	coast:       termbox.Cell{'·', Blue, DarkGreen},
	ocean:       termbox.Cell{'≈', Blue, DarkBlue},
	river:       termbox.Cell{'⩬', Cyan, Blue},
	lake:        termbox.Cell{'⩪', Cyan, DarkCyan},
	haven:       termbox.Cell{'⩫', Cyan, DarkBlue},
	midlands:    termbox.Cell{'⩦', Brown, Green},
	tundra:      termbox.Cell{'⋄', White, Grey},
	forest:      termbox.Cell{'ꙥ', Green, DarkBrown},
	hills:       termbox.Cell{'∿', Green, DarkGreen},
	cliffs:      termbox.Cell{'√', White, DarkGrey},
	highlands:   termbox.Cell{'≏', DarkOrange, DarkGreen},
	mountains:   termbox.Cell{'⩟', Grey, DarkGreen},
	fourteeners: termbox.Cell{'⩓', White, Grey},
}

type manmadeType uint8

const (
	castleM manmadeType = iota
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

type manmadeFeature interface {
	getOwner() *owner
	getTile() *tile
	getType() manmadeType
}

type manmade struct {
	name        string
	owner       *owner
	manmadeType manmadeType
	tile        *tile
}

type structureType uint8

const (
	admin structureType = iota
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
	parchmentMakers  = 419
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
	ropeMakers       = 803
	tanners          = 813
	buckleMakers     = 822
	cutlers          = 831
	fullers          = 840
	harnessMakers    = 849
	painters         = 858
	woodcarvers      = 866
	lampworkers      = 873
	instrumentMakers = 880
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
	sewists          = 992
	alchemists       = 994
	bellMakers       = 996
	dyeMakers        = 998
	inventors        = 1000
)

type shopType int

const (
	outfitters         = 97
	grocers            = 194
	diarysellers       = 270
	launderers         = 346
	prostitutes        = 422
	furrierMerchants   = 498
	tailors            = 558
	barbers            = 607
	drapers            = 656
	flowerSellers      = 705
	jewelers           = 745
	mercers            = 768
	engravers          = 790
	pawnBrokers        = 812
	haberdashers       = 832
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
	grainers           = 996
	tobacconists       = 999
	magicMerchants     = 1000
)

type officeType int

const (
	livestockMerchants = 200
	carpenters         = 360
	masonsO            = 474
	pawnbrokerO        = 546
	wineMerchantsO     = 611
	doctorsUnlicensed  = 661
	woolMerchantsO     = 706
	beerMerchantsO     = 756
	spceMerchants      = 786
	doctorsLicensed    = 815
	copyists           = 842
	moneychangers      = 864
	scholar            = 884
	advocates          = 902
	historians         = 918
	engineers          = 931
	architects         = 941
	astrologers        = 951
	grainMerchantsO    = 961
	tobaccoMerchantsO  = 971
	bankers            = 980
	slavers            = 989
	cartographers      = 997
	magicMerchantsO    = 1000
)

type buildingStyle uint8

const (
	derelictStyle buildingStyle = iota
	basicStyle
	ornateStyle
	imperialStyle
)

type structure struct {
	structureType structureType
	buildingStyle buildingStyle
}

type districtType uint8

const (
	slumDistrict districtType = iota
	gateDistrict
	seaDistrict
	industrial
	riverDistrict
	militaryDistrict
	craftDistrict
	merchantDistrict
	administrativeDistrict
	marketDistrict
	patriciateDistrict
)

type district struct {
	location
	districtType districtType
	blocks       []block
}

type block struct {
	location
	structures []structure
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
	trees, stone, tin, copper, iron, adamantine int
}

type simpleTile struct{ biome }

func (t simpleTile) Cell() termbox.Cell {
	return biomeCells[t.biome]
}

type tile struct {
	naturalResources
	location
	biome
	claimedBy           *owner
	manmadeFeature      *manmadeFeature
	roughness, wildness int
}

func (t tile) Cell() termbox.Cell {
	return biomeCells[t.biome]
}
