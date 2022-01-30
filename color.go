package poker

const (
	Red   Color = "R"
	Black Color = "B"
)

type Color string

func (r Color) String() string {
	return string(r)
}

func ValidateColor(color Color) bool {
	switch color {
	case Red, Black:
		return true
	default:
		return false
	}
}
