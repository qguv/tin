package buildings

type Cardinal uint8

const (
	East Cardinal = iota
	North
	West
	South
)

type Style uint8

const (
	Japanese Style = iota
	Slavic
	W_euro   // western european, e.g. German, French
	E_euro   // eastern european but non-slavic, e.g. Czech
	Baroque  // absurdly adorned
	American // our unique style of old-world architecture
	Mexican  // mayan/aztec step designs
)

type building struct {
	style Style
	gates []Cardinal
	walls []Cardinal
}


