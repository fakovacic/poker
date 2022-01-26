package poker

func IsFlush(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}

	check := make(map[string]int)
	for i := range cards {
		_, ok := check[cards[i].Suite]
		if !ok {
			check[cards[i].Suite] = 1

			continue
		}

		check[cards[i].Suite]++

		if check[cards[i].Suite] == 5 {
			return true
		}
	}

	return false
}
