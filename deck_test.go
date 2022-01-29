package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestDeck(t *testing.T) {
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

	t.Run("deck", func(t *testing.T) {
		deck := poker.NewDeck()

		if len(deck.Cards()) != 52 {
			t.Errorf("expected deck: '%v' got: '%v'", 52, len(deck.Cards()))
		}

		deck.Switch(cards)

		deckCards := deck.Cards()
		if len(deckCards) != 5 {
			t.Errorf("expected deck: '%v' got: '%v'", 5, len(deckCards))
		}

		for i := range deckCards {
			if deckCards[i] != cards[i] {
				t.Errorf("expected deck card %d: '%v' got: '%v'", i, cards[i], deckCards[i])
			}
		}
	})
}

func TestDeckDrawDiscard(t *testing.T) {
	t.Run("deck draw - discard", func(t *testing.T) {
		deck := poker.NewDeck()

		if len(deck.Cards()) != 52 {
			t.Errorf("expected deck: '%v' got: '%v'", 52, len(deck.Cards()))
		}

		deck.Draw()

		if len(deck.Cards()) != 51 {
			t.Errorf("expected deck: '%v' got: '%v'", 51, len(deck.Cards()))
		}

		deck.Discard()

		if len(deck.Cards()) != 50 {
			t.Errorf("expected deck: '%v' got: '%v'", 51, len(deck.Cards()))
		}
	})
}
