package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
)

func TestIsFourOfAKind(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult bool
	}{
		{
			it: "ok, 4 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.King,
				},
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "not ok, 4 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.King,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is.New(t).Equal(tc.expectedResult, poker.IsFourOfAKind(tc.cards))
		})
	}
}
