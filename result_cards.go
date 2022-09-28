package poker

import "sort"

func ResultCards(result Result, cards []*Card) []int64 {
	switch result {
	case RoyalFlush:
		return RoyalFlushCards(cards)
	case StraightFlush:
		return StraightFlushCards(cards)
	case FourOfAKind:
		return FourOfAKindCards(cards)
	case FullHouse:
		return FullHouseCards(cards)
	case Flush:
		return FlushCards(cards)
	case Straight:
		return StraightCards(cards)
	case ThreeOfAKind:
		return ThreeOfAKindCards(cards)
	case TwoPairs:
		return TwoPairsCards(cards)
	case Pair:
		return PairCards(cards)
	case HighCard:
		return HighCardCards(cards)
	default:
		return nil
	}
}

func RoyalFlushCards(cards []*Card) []int64 {
	if len(cards) < 5 {
		return nil
	}

	type royalFlush struct {
		ten   bool
		jack  bool
		queen bool
		king  bool
		ace   bool
	}

	c := make(map[Suite]royalFlush)
	cl := make(map[Suite][]int64)

	for i := range cards {
		suite := cards[i].Suite

		val, ok := c[suite]
		if !ok {
			val = royalFlush{
				ace:   false,
				king:  false,
				queen: false,
				jack:  false,
				ten:   false,
			}
		}

		switch cards[i].Rank {
		case Ace:
			val.ace = true

			cl[suite] = append(cl[suite], int64(i))
		case King:
			val.king = true

			cl[suite] = append(cl[suite], int64(i))
		case Queen:
			val.queen = true

			cl[suite] = append(cl[suite], int64(i))
		case Jack:
			val.jack = true

			cl[suite] = append(cl[suite], int64(i))
		case Ten:
			val.ten = true

			cl[suite] = append(cl[suite], int64(i))

		case Eight, Five, Four, Nine, Seven, Six, Three, Two:
			continue
		}

		c[suite] = val
	}

	for i := range c {
		if c[i].ten && c[i].jack && c[i].queen && c[i].king && c[i].ace {
			return cl[i]
		}
	}

	return nil
}

func StraightFlushCards(cards []*Card) []int64 {
	if len(cards) < 5 {
		return nil
	}

	var flush bool

	type fc struct {
		num   int64
		rank  int64
		suite Suite
	}

	c := make(map[Suite][]fc)

	for i := range cards {
		suite := cards[i].Suite
		c[suite] = append(c[suite], fc{
			num:   int64(i),
			rank:  cards[i].Rank.Score(),
			suite: cards[i].Suite,
		})

		if len(c[cards[i].Suite]) == 5 {
			flush = true
		}
	}

	if !flush {
		return nil
	}

	straightCheck := make(map[Suite][]int64)

	for x := range c {
		if len(c[x]) >= 5 {
			cl := c[x]

			sort.Slice(cl, func(i, j int) bool {
				return cl[i].rank > cl[j].rank
			})

			for y := range cl {
				straightCheck[cl[y].suite] = append(straightCheck[cl[y].suite], cl[y].num)
			}
		}
	}

	for key := range straightCheck {
		var rankKeys []int

		cc := make(map[int]int64)

		for i := range straightCheck[key] {
			cardNum := straightCheck[key][i]

			card := cards[cardNum]

			if card.Rank == Ace {
				cc[ranks[card.Rank+"Low"]] = cardNum
				cc[ranks[card.Rank+"High"]] = cardNum

				rankKeys = append(rankKeys, ranks[card.Rank+"Low"])
				rankKeys = append(rankKeys, ranks[card.Rank+"High"])

				continue
			}

			cc[ranks[card.Rank]] = cardNum

			rankKeys = append(rankKeys, ranks[card.Rank])
		}

		sort.Slice(rankKeys, func(i, j int) bool {
			return rankKeys[i] > rankKeys[j]
		})

		previousCard := 0

		var ck []int64

		for i := range rankKeys {
			if i == 0 {
				previousCard = rankKeys[i]

				ck = append(ck, cc[previousCard])

				continue
			}

			if len(ck) == 5 {
				continue
			}

			if previousCard == -1 {
				previousCard = rankKeys[i-1]

				ck = append(ck, cc[previousCard])
			}

			currentCard := rankKeys[i]
			if (currentCard + 1) == previousCard {
				ck = append(ck, cc[currentCard])
			}

			if (currentCard + 1) != previousCard {
				ck = []int64{}
				previousCard = -1

				continue
			}

			previousCard = rankKeys[i]
		}

		if len(ck) == 5 {
			sort.Slice(ck, func(i, j int) bool {
				return ck[i] < ck[j]
			})

			return ck
		}
	}

	return nil
}

func FourOfAKindCards(cards []*Card) []int64 {
	switch {
	case len(cards) < 4:
		return nil
	case len(cards) == 4:
		return []int64{
			0, 1, 2, 3,
		}
	default:
		var rankKeys []int64

		c := make(map[int64][]int64)

		for i := range cards {
			cs := cards[i].Rank.Score()

			_, ok := c[cs]
			if !ok {
				c[cs] = append(c[cs], int64(i))

				rankKeys = append(rankKeys, cs)

				continue
			}

			c[cs] = append(c[cs], int64(i))
		}

		sort.Slice(rankKeys, func(i, j int) bool {
			return rankKeys[i] < rankKeys[j]
		})

		for i := range rankKeys {
			if len(c[rankKeys[i]]) == 4 {
				return c[rankKeys[i]]
			}
		}

		return nil
	}
}

func FullHouseCards(cards []*Card) []int64 {
	if len(cards) < 5 {
		return nil
	}

	var (
		pairs      []int64
		treeOfKind []int64
	)

	var rankKeys []int64

	c := make(map[int64][]int64)

	for i := range cards {
		cs := cards[i].Rank.Score()

		_, ok := c[cs]
		if !ok {
			c[cs] = append(c[cs], int64(i))

			rankKeys = append(rankKeys, cs)

			continue
		}

		c[cs] = append(c[cs], int64(i))
	}

	sort.Slice(rankKeys, func(i, j int) bool {
		return rankKeys[i] > rankKeys[j]
	})

	for i := range rankKeys {
		if len(c[rankKeys[i]]) >= 3 && len(treeOfKind) == 0 {
			treeOfKind = c[rankKeys[i]]

			continue
		}

		if len(c[rankKeys[i]]) >= 2 && len(pairs) == 0 {
			pairs = c[rankKeys[i]][:2]

			continue
		}
	}

	if len(pairs) != 0 && len(treeOfKind) != 0 {
		ci := make([]int64, 5)
		ci[0] = pairs[0]
		ci[1] = pairs[1]
		ci[2] = treeOfKind[0]
		ci[3] = treeOfKind[1]
		ci[4] = treeOfKind[2]

		sort.Slice(ci, func(i, j int) bool {
			return ci[i] < ci[j]
		})

		return ci
	}

	return nil
}

func FlushCards(cards []*Card) []int64 {
	if len(cards) < 5 {
		return nil
	}

	var flush bool

	type fc struct {
		num  int64
		rank int64
	}

	c := make(map[Suite][]fc)

	for i := range cards {
		suite := cards[i].Suite
		c[suite] = append(c[suite], fc{
			num:  int64(i),
			rank: cards[i].Rank.Score(),
		})

		if len(c[cards[i].Suite]) == 5 {
			flush = true
		}
	}

	if !flush {
		return nil
	}

	for x := range c {
		if len(c[x]) >= 5 {
			cl := c[x]

			sort.Slice(cl, func(i, j int) bool {
				return cl[i].rank > cl[j].rank
			})

			var ck []int64
			for y := range cl {
				ck = append(ck, cl[y].num)
				if len(ck) == 5 {
					sort.Slice(ck, func(i, j int) bool {
						return ck[i] < ck[j]
					})

					return ck
				}
			}
		}
	}

	return nil
}

func StraightCards(cards []*Card) []int64 {
	if len(cards) < 5 {
		return nil
	}

	rankKeys := make([]int, 0, 6)

	c := make(map[int]int64)

	for i := range cards {
		if cards[i].Rank == Ace {
			c[ranks[cards[i].Rank+"Low"]] = int64(i)
			c[ranks[cards[i].Rank+"High"]] = int64(i)

			rankKeys = append(rankKeys, ranks[cards[i].Rank+"Low"])
			rankKeys = append(rankKeys, ranks[cards[i].Rank+"High"])

			continue
		}

		c[ranks[cards[i].Rank]] = int64(i)
		rankKeys = append(rankKeys, ranks[cards[i].Rank])
	}

	sort.Slice(rankKeys, func(i, j int) bool {
		return rankKeys[i] > rankKeys[j]
	})

	previousCard := 0

	var ck []int64

	for i := range rankKeys {
		if i == 0 {
			previousCard = rankKeys[i]

			ck = append(ck, c[previousCard])

			continue
		}

		if len(ck) == 5 {
			continue
		}

		if previousCard == -1 {
			previousCard = rankKeys[i-1]

			ck = append(ck, c[previousCard])
		}

		currentCard := rankKeys[i]
		if (currentCard + 1) == previousCard {
			ck = append(ck, c[currentCard])
		}

		if (currentCard + 1) != previousCard {
			ck = []int64{}
			previousCard = -1

			continue
		}

		previousCard = rankKeys[i]
	}

	if len(ck) == 5 {
		sort.Slice(ck, func(i, j int) bool {
			return ck[i] < ck[j]
		})

		return ck
	}

	return nil
}

func ThreeOfAKindCards(cards []*Card) []int64 {
	switch {
	case len(cards) < 3:
		return nil
	case len(cards) == 3:
		return []int64{
			0, 1, 2,
		}
	default:
		var rankKeys []int64

		c := make(map[int64][]int64)

		for i := range cards {
			cs := cards[i].Rank.Score()

			_, ok := c[cs]
			if !ok {
				c[cs] = append(c[cs], int64(i))

				rankKeys = append(rankKeys, cs)

				continue
			}

			c[cs] = append(c[cs], int64(i))
		}

		sort.Slice(rankKeys, func(i, j int) bool {
			return rankKeys[i] < rankKeys[j]
		})

		for i := range rankKeys {
			if len(c[rankKeys[i]]) == 3 {
				return c[rankKeys[i]]
			}
		}

		return nil
	}
}

func TwoPairsCards(cards []*Card) []int64 {
	switch {
	case len(cards) < 4:
		return nil
	case len(cards) == 4:
		return []int64{
			0, 1, 2, 3,
		}
	default:
		var rankKeys []int64

		c := make(map[int64][]int64)

		for i := range cards {
			cs := cards[i].Rank.Score()

			_, ok := c[cs]
			if !ok {
				c[cs] = append(c[cs], int64(i))

				rankKeys = append(rankKeys, cs)

				continue
			}

			c[cs] = append(c[cs], int64(i))
		}

		sort.Slice(rankKeys, func(i, j int) bool {
			return rankKeys[i] < rankKeys[j]
		})

		var pairs []int64

		for i := range rankKeys {
			if len(c[rankKeys[i]]) == 2 {
				pairs = append(pairs, c[rankKeys[i]]...)
			}

			if len(pairs) == 4 {
				sort.Slice(pairs, func(i, j int) bool {
					return pairs[i] < pairs[j]
				})

				return pairs
			}
		}

		return nil
	}
}

func PairCards(cards []*Card) []int64 {
	switch {
	case len(cards) < 2:
		return nil
	case len(cards) == 2:
		return []int64{
			0, 1,
		}
	default:
		var rankKeys []int

		c := make(map[int64][]int64)

		for i := range cards {
			cs := cards[i].Rank.Score()

			_, ok := c[cs]
			if !ok {
				c[cs] = append(c[cs], int64(i))

				rankKeys = append(rankKeys, int(cs))

				continue
			}

			c[cs] = append(c[cs], int64(i))
		}

		sort.Slice(rankKeys, func(i, j int) bool {
			return rankKeys[i] > rankKeys[j]
		})

		for i := range rankKeys {
			if len(c[int64(rankKeys[i])]) == 2 {
				return c[int64(rankKeys[i])]
			}
		}

		return nil
	}
}

func HighCardCards(cards []*Card) []int64 {
	switch {
	case len(cards) < 1:
		return nil
	case len(cards) == 1:
		return []int64{
			0,
		}
	default:
		var rankKeys []int

		c := make(map[int64]int64)

		for i := range cards {
			cs := cards[i].Rank.Score()
			rankKeys = append(rankKeys, int(cs))
			c[cs] = int64(i)
		}

		sort.Slice(rankKeys, func(i, j int) bool {
			return rankKeys[i] > rankKeys[j]
		})

		return []int64{c[int64(rankKeys[0])]}
	}
}
