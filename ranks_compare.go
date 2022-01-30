package poker

const (
	High  string = "high"
	Low   string = "low"
	Equal string = "equal"
)

func CompareCardsRank(xCard, yCard *Card) string {
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
