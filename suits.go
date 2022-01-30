package poker

const (
	Clubs    Suite = "C"
	Diamonds Suite = "D"
	Hearts   Suite = "H"
	Spades   Suite = "S"
)

type Suite string

func (r Suite) String() string {
	return string(r)
}

func (r Suite) Color() Color {
	switch r {
	case Clubs:
		return Black
	case Spades:
		return Black
	case Diamonds:
		return Red
	case Hearts:
		return Red
	default:
		return ""
	}
}

var suites = []Suite{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}

func ValidateSuite(suite Suite) bool {
	switch suite {
	case Clubs, Diamonds, Hearts, Spades:
		return true
	default:
		return false
	}
}
