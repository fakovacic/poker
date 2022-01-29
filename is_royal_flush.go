package poker

func IsRoyalFlush(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}

	type royalFlush struct {
		ten   bool
		jack  bool
		queen bool
		king  bool
		ace   bool
	}

	c := make(map[string]royalFlush)

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
		case King:
			val.king = true
		case Queen:
			val.queen = true
		case Jack:
			val.jack = true
		case Ten:
			val.ten = true
		}

		c[suite] = val
	}

	for i := range c {
		if c[i].ten && c[i].jack && c[i].queen && c[i].king && c[i].ace {
			return true
		}
	}

	return false
}
