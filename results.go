package poker

const (
	HighCard      Result = "high-card"
	Pair          Result = "pair"
	TwoPairs      Result = "two-pairs"
	ThreeOfAKind  Result = "three-of-a-kind"
	Straight      Result = "straight"
	Flush         Result = "flush"
	FullHouse     Result = "full-house"
	FourOfAKind   Result = "four-of-a-kind"
	StraightFlush Result = "straight-flush"
	RoyalFlush    Result = "royal-flush"
)

type Result string

func (r Result) String() string {
	return string(r)
}

func (r Result) Score() int64 {
	return resultScore[r]
}

func ParseResult(result string) (Result, bool) {
	switch result {
	case "high-card":
		return HighCard, true
	case "pair":
		return Pair, true
	case "two-pairs":
		return TwoPairs, true
	case "three-of-a-kind":
		return ThreeOfAKind, true
	case "straight":
		return Straight, true
	case "flush":
		return Flush, true
	case "full-house":
		return FullHouse, true
	case "four-of-a-kind":
		return FourOfAKind, true
	case "straight-flush":
		return StraightFlush, true
	case "royal-flush":
		return RoyalFlush, true
	default:
		return "", false
	}
}

var resultScore = map[Result]int64{
	HighCard:      1,
	Pair:          3,
	TwoPairs:      4,
	ThreeOfAKind:  5,
	Straight:      6,
	Flush:         7,
	FullHouse:     8,
	FourOfAKind:   9,
	StraightFlush: 10,
	RoyalFlush:    11,
}

var ResultsList = []Result{
	RoyalFlush,
	StraightFlush,
	FourOfAKind,
	FullHouse,
	Flush,
	Straight,
	ThreeOfAKind,
	TwoPairs,
	Pair,
	HighCard,
}
