package poker

func IsFullHouse(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}

	var (
		pairs      bool
		treeOfKind bool
	)

	check := make(map[Rank]int, len(cards))
	for i := range cards {
		_, ok := check[cards[i].Rank]
		if !ok {
			check[cards[i].Rank] = 1

			continue
		}

		check[cards[i].Rank]++
	}

	for deck := range check {
		if check[deck] == 2 {
			pairs = true
		}

		if check[deck] == 3 {
			treeOfKind = true
		}
	}

	if pairs && treeOfKind {
		return true
	}

	return false
}
