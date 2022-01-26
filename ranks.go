package poker

const (
	Ace   string = "A"
	King  string = "K"
	Queen string = "Q"
	Jack  string = "J"
	Ten   string = "10"
	Nine  string = "9"
	Eight string = "8"
	Seven string = "7"
	Six   string = "6"
	Five  string = "5"
	Four  string = "4"
	Three string = "3"
	Two   string = "2"
)

var ranksList = []string{
	Ace,
	King,
	Queen,
	Jack,
	Ten,
	Nine,
	Eight,
	Seven,
	Six,
	Five,
	Four,
	Three,
	Two,
}

var ranks = map[string]int{
	Ace + "Low":  1,
	Two:          2,
	Three:        3,
	Four:         4,
	Five:         5,
	Six:          6,
	Seven:        7,
	Eight:        8,
	Nine:         9,
	Ten:          10,
	Jack:         11,
	Queen:        12,
	King:         13,
	Ace + "High": 14,
}

var ranksScore = map[string]int64{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}
