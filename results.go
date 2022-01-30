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

func (r Result) Score() int64 {
	return resultScore[r]
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
