package poker

func IsFlush(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}

	check := make(map[Suite]int, 4)
	for i := range cards {
		val, ok := check[cards[i].Suite]
		if !ok {
			check[cards[i].Suite] = 1

			continue
		}

		val++

		if val == 5 {
			return true
		}

		check[cards[i].Suite] = val
	}

	return false
}
