package poker

func IsFourOfAKind(cards []*Card) bool {
	if len(cards) < 4 {
		return false
	}

	check := make(map[Rank]int, len(cards))
	for i := range cards {
		val, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		val++

		if val == 4 {
			return true
		}

		check[cards[i].Rank] = val
	}

	return false
}
