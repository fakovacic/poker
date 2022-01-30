package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestCardsScore(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		scoringCards   []int64
		expectedResult int64
	}{
		{
			it: "high-card",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
			},
			scoringCards: []int64{
				4,
			},
			expectedResult: 14,
		},
		{
			it: "pair",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Two,
				},
			},
			scoringCards: []int64{
				0, 1,
			},
			expectedResult: 28,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.CardsScore(tc.cards, tc.scoringCards)
			if tc.expectedResult != res {
				t.Errorf("expected: '%d' got: '%d'", tc.expectedResult, res)
			}
		})
	}
}
