package poker

import (
	"sort"
)

func IsStraight(cards []*Card) bool {
	if len(cards) != 5 {
		return false
	}

	check := make([]int, 0)

	for i := range cards {
		if cards[i].Rank == Ace {
			check = append(check, ranks[cards[i].Rank+"Low"])
			check = append(check, ranks[cards[i].Rank+"High"])

			continue
		}

		check = append(check, ranks[cards[i].Rank])
	}

	if len(check) < 5 || len(check) > 6 {
		return false
	}

	sort.Ints(check)

	previousCard := 0
	ok := 0

	for i := range check {
		if i == 0 {
			previousCard = check[i]

			continue
		}

		currentCard := check[i]
		if (currentCard - 1) == previousCard {
			ok++
		}

		if ok == 4 {
			return true
		}

		if (currentCard - 1) != previousCard {
			ok = 0
		}

		previousCard = check[i]
	}

	return ok == 4
}
