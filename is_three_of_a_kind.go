package poker

func IsThreeOfAKind(cards []*Card) bool {
	if len(cards) < 3 {
		return false
	}

	check := make(map[Rank]int, len(cards))

	for i := range cards {
		r := cards[i].Rank

		val, ok := check[r]
		if !ok {
			check[r] = 1

			continue
		}

		val++

		if val == 3 {
			return true
		}

		check[r] = val
	}

	return false
}
