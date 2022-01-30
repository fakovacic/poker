package poker

func CardsResult(cards []*Card) Result {
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
