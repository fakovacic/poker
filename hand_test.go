package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestHand(t *testing.T) {
	cards := []*poker.Card{
		{
			Suite: poker.Spades,
			Rank:  poker.Ten,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.Jack,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.Queen,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.King,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.Two,
		},
	}

	t.Run("hand", func(t *testing.T) {
		hand := poker.NewHand()

		if len(hand.Cards()) != 0 {
			t.Errorf("expected hand: '%v' got: '%v'", 0, len(hand.Cards()))
		}

		for i := range cards {
			hand.Deal(cards[i])
		}

		if len(hand.Cards()) != 5 {
			t.Errorf("expected hand: '%v' got: '%v'", 5, len(hand.Cards()))
		}
	})
}
