package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
)

func TestIsStraight(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult bool
	}{
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Two,
				},
				{
					Rank: poker.Three,
				},
				{
					Rank: poker.Four,
				},
				{
					Rank: poker.Five,
				},
				{
					Rank: poker.Six,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Two,
				},
				{
					Rank: poker.Six,
				},
				{
					Rank: poker.Four,
				},
				{
					Rank: poker.Five,
				},
				{
					Rank: poker.Three,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Nine,
				},
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Jack,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Jack,
				},
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Nine,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: true,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Nine,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Six,
				},
				{
					Rank: poker.Nine,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Two,
				},
				{
					Rank: poker.Six,
				},
				{
					Rank: poker.Nine,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 4 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Jack,
				},
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Nine,
				},
				{
					Rank: poker.Queen,
				},
			},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is.New(t).Equal(tc.expectedResult, poker.IsStraight(tc.cards))
		})
	}

}

func TestIsStraightWithAce(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult bool
	}{
		{
			it: "ok, 5 cards low",
			cards: []*poker.Card{
				{
					Rank: poker.Two,
				},
				{
					Rank: poker.Three,
				},
				{
					Rank: poker.Four,
				},
				{
					Rank: poker.Five,
				},
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards high",
			cards: []*poker.Card{
				{
					Rank: poker.Ten,
				},
				{
					Rank: poker.Jack,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "not ok, 5 cards high",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Two,
				},
				{
					Rank: poker.Jack,
				},
				{
					Rank: poker.Queen,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards low",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.Two,
				},
				{
					Rank: poker.Three,
				},
				{
					Rank: poker.King,
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
			is.New(t).Equal(tc.expectedResult, poker.IsStraight(tc.cards))
		})
	}
}
