package poker

func IsPair(cards []*Card) bool {
	if len(cards) < 2 {
		return false
	}

	check := make(map[Rank]int, len(cards))
	for _, c := range cards {
		_, ok := check[c.Rank]
		if !ok {
			check[c.Rank] = 1

			continue
		}

		if ok {
			return true
		}
	}

	return false
}
