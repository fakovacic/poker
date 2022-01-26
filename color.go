package poker

const (
	Red   string = "R"
	Black string = "B"
)

func Color(color string) string {
	switch color {
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

func ValidateColor(color string) bool {
	switch color {
	case Red, Black:
		return true
	default:
		return false
	}
}
