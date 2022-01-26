package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
)

func TestOddsHand(t *testing.T) {
	cases := []struct {
		it                   string
		hand                 []*poker.Card
		deck                 []*poker.Card
		expectedCombinations int64
		expectedOdds         map[string]float64
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
			expectedOdds: map[string]float64{
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
			expectedOdds: map[string]float64{
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
			expectedOdds: map[string]float64{
				poker.FourOfAKind: 100,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			combs, odds := poker.OddsHand(tc.hand, tc.deck)

			checkIs.Equal(tc.expectedCombinations, combs)
			checkIs.Equal(tc.expectedOdds, odds)
		})
	}
}
