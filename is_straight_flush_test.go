package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestIsStraightFlush(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult bool
	}{
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Three,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Hearts,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Three,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Nine,
				},
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
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Clubs,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Nine,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.King,
				},
			},
			expectedResult: true,
		},

		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Nine,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Nine,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Nine,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Three,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.IsStraightFlush(tc.cards)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
