package poker

func IsPair(cards []*Card) bool {
	if len(cards) < 2 {
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

		if val == 2 {
			return true
		}

		check[cards[i].Rank] = val
	}

	return false
}
