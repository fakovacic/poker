package poker

const (
	Clubs    string = "C"
	Diamonds string = "D"
	Hearts   string = "H"
	Spades   string = "S"
)

var suites = []string{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}

func ValidateSuite(suite string) bool {
	switch suite {
	case Clubs, Diamonds, Hearts, Spades:
		return true
	default:
		return false
	}
}
