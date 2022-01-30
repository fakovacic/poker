package poker

func CompareCardsRank(xCard, yCard *Card) Compare {
	switch {
	case xCard.Rank.Score() < yCard.Rank.Score():
		return High
	case xCard.Rank.Score() > yCard.Rank.Score():
		return Low
	case xCard.Rank.Score() == yCard.Rank.Score():
		return Equal
	default:
		return ""
	}
}
