package poker

const (
	High  string = "high"
	Low   string = "low"
	Equal string = "equal"
)

func CompareCardsRank(xCard, yCard *Card) string {
	switch {
	case xCard.RankScore() < yCard.RankScore():
		return High
	case xCard.RankScore() > yCard.RankScore():
		return Low
	case xCard.RankScore() == yCard.RankScore():
		return Equal
	default:
		return ""
	}
}
