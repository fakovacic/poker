package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestOddsHand(t *testing.T) {
	cases := []struct {
		it                   string
		hand                 []*poker.Card
		deck                 []*poker.Card
		expectedCombinations int64
		expectedOdds         map[poker.Result]float64
	}{
		{
			it: "pair - 3 cards",
			hand: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Clubs,
				},
				{
					Rank:  poker.King,
					Suite: poker.Diamonds,
				},
			},
			deck: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Hearts,
				},
				{
					Rank:  poker.Queen,
					Suite: poker.Spades,
				},
				{
					Rank:  poker.King,
					Suite: poker.Clubs,
				},
			},
			expectedCombinations: 1,
			expectedOdds: map[poker.Result]float64{
				poker.TwoPairs: 100,
			},
		},
		{
			it: "pair - 5 cards",
			hand: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Clubs,
				},
				{
					Rank:  poker.King,
					Suite: poker.Diamonds,
				},
			},
			deck: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Hearts,
				},
				{
					Rank:  poker.Queen,
					Suite: poker.Spades,
				},
				{
					Rank:  poker.King,
					Suite: poker.Clubs,
				},
				{
					Rank:  poker.Seven,
					Suite: poker.Clubs,
				},
				{
					Rank:  poker.Two,
					Suite: poker.Clubs,
				},
			},
			expectedCombinations: 10,
			expectedOdds: map[poker.Result]float64{
				poker.HighCard: 10,
				poker.Pair:     60,
				poker.TwoPairs: 30,
			},
		},
		{
			it: "four-of-a-kind",
			hand: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Clubs,
				},
				{
					Rank:  poker.Ace,
					Suite: poker.Diamonds,
				},
			},
			deck: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Hearts,
				},
				{
					Rank:  poker.Ace,
					Suite: poker.Spades,
				},
				{
					Rank:  poker.King,
					Suite: poker.Clubs,
				},
			},
			expectedCombinations: 1,
			expectedOdds: map[poker.Result]float64{
				poker.FourOfAKind: 100,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			combs, odds := poker.OddsHand(tc.hand, tc.deck)

			if tc.expectedCombinations != combs {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedCombinations, combs)
			}

			for i := range tc.expectedOdds {
				if tc.expectedOdds[i] != odds[i] {
					t.Errorf("expected odds %s: '%v' got: '%v'", i, tc.expectedOdds[i], odds[i])
				}
			}
		})
	}
}

func TestOddsHandDeck(t *testing.T) {
	cases := []struct {
		it                   string
		hand                 []*poker.Card
		expectedCombinations int64
		expectedOdds         map[poker.Result]float64
	}{
		{
			it: "check",
			hand: []*poker.Card{
				{
					Rank:  poker.Ace,
					Suite: poker.Clubs,
				},
				{
					Rank:  poker.King,
					Suite: poker.Diamonds,
				},
			},
			expectedCombinations: 1,
			expectedOdds: map[poker.Result]float64{
				poker.TwoPairs: 100,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			deck := poker.NewDeck()
			deck.RemoveCards(tc.hand)

			combs, odds := poker.OddsHand(tc.hand, deck.Cards())

			if tc.expectedCombinations != combs {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedCombinations, combs)
			}

			for i := range tc.expectedOdds {
				if tc.expectedOdds[i] != odds[i] {
					t.Errorf("expected odds %s: '%v' got: '%v'", i, tc.expectedOdds[i], odds[i])
				}
			}
		})
	}
}
