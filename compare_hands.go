package poker

import "sort"

type handScore struct {
	key         int64
	resultScore int64
	handScore   int64
}

func CompareHands(hands ...[]*Card) []int64 {
	count := len(hands)

	list := make([]handScore, count)
	for i := range hands {
		result, cards := BestResult(hands[i])

		cardsScore := CardsScore(hands[i], cards)

		list[i] = handScore{
			key:         int64(i),
			handScore:   cardsScore,
			resultScore: result.Score(),
		}
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].resultScore == list[j].resultScore {
			return list[i].handScore > list[j].handScore
		}

		return list[i].resultScore > list[j].resultScore
	})

	res := make([]int64, count)
	for i := range list {
		res[i] = list[i].key
	}

	return res
}
