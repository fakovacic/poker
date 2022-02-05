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

func ParseSuite(suite string) (Suite, bool) {
	switch suite {
	case "C":
		return Clubs, true
	case "D":
		return Diamonds, true
	case "H":
		return Hearts, true
	case "S":
		return Spades, true
	default:
		return "", false
	}
}

var suites = []Suite{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}
