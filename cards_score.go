package poker

func CardsScore(cards []*Card, scoringCards []int64) int64 {
	var score int64
	for i := range scoringCards {
		score += cards[scoringCards[i]].Rank.Score()
	}

	return score
}
