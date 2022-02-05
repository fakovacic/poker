package poker

const (
	Ace   Rank = "A"
	King  Rank = "K"
	Queen Rank = "Q"
	Jack  Rank = "J"
	Ten   Rank = "10"
	Nine  Rank = "9"
	Eight Rank = "8"
	Seven Rank = "7"
	Six   Rank = "6"
	Five  Rank = "5"
	Four  Rank = "4"
	Three Rank = "3"
	Two   Rank = "2"
)

type Rank string

func (r Rank) String() string {
	return string(r)
}

func (r Rank) Score() int64 {
	return ranksScore[r]
}

func (r Rank) ScorePosition() (int64, int64) {
	if r == "" {
		return 0, 0
	}

	rankNum := r.Score()

	var (
		less int64
		more int64
	)

	for k := range ranksScore {
		switch {
		case rankNum < ranksScore[k]:
			more++
		case rankNum > ranksScore[k]:
			less++
		}
	}

	return less, more
}

func ParseRank(rank string) (Rank, bool) {
	switch rank {
	case "A":
		return Ace, true
	case "K":
		return King, true
	case "Q":
		return Queen, true
	case "J":
		return Jack, true
	case "10":
		return Ten, true
	case "9":
		return Nine, true
	case "8":
		return Eight, true
	case "7":
		return Seven, true
	case "6":
		return Six, true
	case "5":
		return Five, true
	case "4":
		return Four, true
	case "3":
		return Three, true
	case "2":
		return Two, true
	default:
		return "", false
	}
}

var ranksScore = map[Rank]int64{
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

var ranksList = []Rank{
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

var ranks = map[Rank]int{
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
