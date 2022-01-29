package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestOddsDeckAllCards(t *testing.T) {
	expectedCombinations := int64(2598960)
	expectedOdds := map[string]float64{
		poker.HighCard:      50.11774,
		poker.Pair:          42.2569,
		poker.TwoPairs:      4.7539,
		poker.ThreeOfAKind:  2.11285,
		poker.Straight:      0.39246,
		poker.Flush:         0.19654,
		poker.FullHouse:     0.14406,
		poker.FourOfAKind:   0.02401,
		poker.StraightFlush: 0.00139,
		poker.RoyalFlush:    0.00015,
	}

	t.Run("all_hands", func(t *testing.T) {
		deck := poker.NewDeck()

		combs, odds := poker.OddsDeck(deck.Cards())

		if expectedCombinations != combs {
			t.Errorf("expected: '%v' got: '%v'", expectedCombinations, combs)
		}

		for i := range expectedOdds {
			if expectedOdds[i] != odds[i] {
				t.Errorf("expected odds %s: '%v' got: '%v'", i, expectedOdds[i], odds[i])
			}
		}
	})
}
