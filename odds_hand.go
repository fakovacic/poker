package poker

import "math"

func OddsHand(hand, deck []*Card) (int64, map[string]float64) {
	s := oddsHand{
		odds:         make(map[string]float64),
		combinations: 0,
		out:          make([]*Card, 0),
	}

	s.oddsCalculateHand(hand, deck, 5-len(hand), 0)

	for key, val := range s.odds {
		s.odds[key] = math.Round(((val/float64(s.combinations))*100)*oddsRounding) / oddsRounding
	}

	return s.combinations, s.odds
}

type oddsHand struct {
	out          []*Card
	combinations int64
	odds         map[string]float64
}

func (s *oddsHand) oddsCalculateHand(hand, deck []*Card, level, i int) {
	if len(deck) == 0 || level > len(deck) {
		return
	}

	if level == 0 {
		c := hand
		c = append(c, s.out...)
		s.combinations++
		s.checkCombination(c)

		return
	}

	var j int
	for j = i; j < len(deck); j++ {
		s.out = append(s.out, deck[j])
		s.oddsCalculateHand(hand, deck, level-1, j+1)
		s.out = s.out[:len(s.out)-1]

		if (len(s.out) - j) > level {
			return
		}
	}
}

func (s *oddsHand) checkCombination(cards []*Card) {
	res := Result(cards)

	s.odds[res]++
}
