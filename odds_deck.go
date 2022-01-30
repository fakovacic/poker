package poker

import (
	"math"
)

type oddsDeck struct {
	out          []*Card
	combinations int64
	odds         map[Result]float64
}

func OddsDeck(deck []*Card) (int64, map[Result]float64) {
	s := oddsDeck{
		odds:         make(map[Result]float64),
		combinations: 0,
		out:          make([]*Card, 0),
	}

	s.oddsCalculateDeck(deck, 5, 0)

	for key, val := range s.odds {
		s.odds[key] = math.Round(((val/float64(s.combinations))*100)*oddsRounding) / oddsRounding
	}

	return s.combinations, s.odds
}

func (s *oddsDeck) oddsCalculateDeck(deck []*Card, level, i int) {
	if len(deck) == 0 || level > len(deck) {
		return
	}

	if level == 0 {
		s.combinations++
		s.checkCombination(s.out)

		return
	}

	var j int
	for j = i; j < len(deck); j++ {
		s.out = append(s.out, deck[j])
		s.oddsCalculateDeck(deck, level-1, j+1)
		s.out = s.out[:len(s.out)-1]

		if (len(s.out) - j) > level {
			return
		}
	}
}

func (s *oddsDeck) checkCombination(cards []*Card) {
	res := CardsResult(cards)

	s.odds[res]++
}
