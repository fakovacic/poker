package poker

func IsTwoPairs(cards []*Card) bool {
	if len(cards) < 4 {
		return false
	}

	var pairs int

	check := make(map[string]int)
	for i := range cards {
		_, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		check[cards[i].Rank]++

		if check[cards[i].Rank] == 2 {
			pairs++
		}

		if pairs == 2 {
			return true
		}
	}

	return false
}
