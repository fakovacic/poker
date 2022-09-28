package poker

import "sort"

func IsStraightFlush(cards []*Card) bool {
	if len(cards) != 5 {
		return false
	}

	var flush bool

	checkFlush := make(map[Suite]int, 4)
	checkstraight := make([]int, 0, 6)

	for i := range cards {
		val, ok := checkFlush[cards[i].Suite]
		if !ok {
			val = 1
		} else {
			val++
		}

		checkFlush[cards[i].Suite] = val

		if cards[i].Rank == Ace {
			checkstraight = append(checkstraight, ranks[cards[i].Rank+"Low"])
			checkstraight = append(checkstraight, ranks[cards[i].Rank+"High"])

			continue
		}

		checkstraight = append(checkstraight, ranks[cards[i].Rank])

		if val == 5 {
			flush = true
		}
	}

	if !flush {
		return false
	}

	if len(checkstraight) < 5 || len(checkstraight) > 6 {
		return false
	}

	sort.Ints(checkstraight)

	previousCard := 0
	ok := 0

	for i := range checkstraight {
		if i == 0 {
			previousCard = checkstraight[i]

			continue
		}

		currentCard := checkstraight[i]
		if (currentCard - 1) == previousCard {
			ok++
		}

		if ok == 4 {
			return true
		}

		if (currentCard - 1) != previousCard {
			ok = 0
		}

		previousCard = checkstraight[i]
	}

	return ok == 4
}
