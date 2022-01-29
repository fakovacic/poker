package poker

import "sort"

func IsStraightFlush(cards []*Card) bool {
	if len(cards) != 5 {
		return false
	}

	var flush bool

	checkFlush := make(map[string]int)
	checkstraight := make([]int, 0)

	for i := range cards {
		_, ok := checkFlush[cards[i].Suite]
		if !ok {
			checkFlush[cards[i].Suite] = 1
		} else {
			checkFlush[cards[i].Suite]++
		}

		if cards[i].Rank == Ace {
			checkstraight = append(checkstraight, ranks[cards[i].Rank+"Low"])
			checkstraight = append(checkstraight, ranks[cards[i].Rank+"High"])

			continue
		}

		checkstraight = append(checkstraight, ranks[cards[i].Rank])

		if checkFlush[cards[i].Suite] == 5 {
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
