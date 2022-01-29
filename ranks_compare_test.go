package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestCompareCardsRank(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult string
	}{
		{
			it: "high",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.High,
		},
		{
			it: "low",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Nine,
				},
			},
			expectedResult: poker.Low,
		},
		{
			it: "equal",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ten,
				},
			},
			expectedResult: poker.Equal,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.CompareCardsRank(tc.cards[0], tc.cards[1])
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
