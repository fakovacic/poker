package poker

const (
	Red   Color = "R"
	Black Color = "B"
)

type Color string

func (r Color) String() string {
	return string(r)
}

func ParseColor(color string) (Color, bool) {
	switch color {
	case "R":
		return Red, true
	case "B":
		return Black, true
	default:
		return "", false
	}
}
