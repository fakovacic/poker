package poker

const (
	HighCard      string = "high-card"
	Pair          string = "pair"
	TwoPairs      string = "two-pairs"
	ThreeOfAKind  string = "three-of-a-kind"
	Straight      string = "straight"
	Flush         string = "flush"
	FullHouse     string = "full-house"
	FourOfAKind   string = "four-of-a-kind"
	StraightFlush string = "straight-flush"
	RoyalFlush    string = "royal-flush"
)

var ResultsList = []string{
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

var resultScore = map[string]int64{
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

func Result(cards []*Card) string {
	var ok bool

	ok = IsRoyalFlush(cards)
	if ok {
		return RoyalFlush
	}

	ok = IsStraightFlush(cards)
	if ok {
		return StraightFlush
	}

	ok = IsFourOfAKind(cards)
	if ok {
		return FourOfAKind
	}

	ok = IsFullHouse(cards)
	if ok {
		return FullHouse
	}

	ok = IsFlush(cards)
	if ok {
		return Flush
	}

	ok = IsStraight(cards)
	if ok {
		return Straight
	}

	ok = IsThreeOfAKind(cards)
	if ok {
		return ThreeOfAKind
	}

	ok = IsTwoPairs(cards)
	if ok {
		return TwoPairs
	}

	ok = IsPair(cards)
	if ok {
		return Pair
	}

	return HighCard
}
