package poker

func IsThreeOfAKind(cards []*Card) bool {
	if len(cards) < 3 {
		return false
	}

	check := make(map[string]int)

	for i := range cards {
		r := cards[i].Rank

		_, ok := check[r]
		if !ok {
			check[r] = 1

			continue
		}

		check[r]++

		if check[r] == 3 {
			return true
		}
	}

	return false
}
