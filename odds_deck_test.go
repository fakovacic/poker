package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
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
		checkIs := is.New(t)

		deck := poker.NewDeck()

		combs, odds := poker.OddsDeck(deck.Cards())

		checkIs.Equal(expectedCombinations, combs)
		checkIs.Equal(expectedOdds, odds)
	})
}
