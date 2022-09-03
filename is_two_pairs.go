package poker

func IsTwoPairs(cards []*Card) bool {
	if len(cards) < 4 {
		return false
	}

	var pairs int

	check := make(map[Rank]int, len(cards))
	for i := range cards {
		val, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		val++

		if val == 2 {
			pairs++
		}

		if pairs == 2 {
			return true
		}

		check[cards[i].Rank] = val
	}

	return false
}
