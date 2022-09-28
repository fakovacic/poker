package poker

func IsThreeOfAKind(cards []*Card) bool {
	if len(cards) < 3 {
		return false
	}

	check := make(map[Rank]int, len(cards))

	for i := range cards {
		_, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		check[cards[i].Rank]++

		if check[cards[i].Rank] == 3 {
			return true
		}
	}

	return false
}
