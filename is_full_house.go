package poker

func IsFullHouse(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}

	var (
		pairs      Rank
		treeOfKind Rank
	)

	check := make(map[Rank]int)
	for i := range cards {
		_, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		check[cards[i].Rank]++
	}

	for deck, num := range check {
		if num == 2 {
			pairs = deck
		}

		if num == 3 {
			treeOfKind = deck
		}
	}

	if pairs != "" && treeOfKind != "" {
		return true
	}

	return false
}
