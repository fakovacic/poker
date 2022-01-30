package poker

func IsFourOfAKind(cards []*Card) bool {
	if len(cards) < 4 {
		return false
	}

	check := make(map[Rank]int)
	for i := range cards {
		_, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		check[cards[i].Rank]++

		if check[cards[i].Rank] == 4 {
			return true
		}
	}

	return false
}
