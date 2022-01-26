package poker

func IsPair(cards []*Card) bool {
	if len(cards) < 2 {
		return false
	}

	check := make(map[string]int)
	for i := range cards {
		_, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		check[cards[i].Rank]++

		if check[cards[i].Rank] == 2 {
			return true
		}
	}

	return false
}
