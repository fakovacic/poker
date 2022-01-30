package poker

const (
	High  Compare = "high"
	Low   Compare = "low"
	Equal Compare = "equal"
)

type Compare string

func (c Compare) String() string {
	return string(c)
}
