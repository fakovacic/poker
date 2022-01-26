package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
)

func TestIsRoyalFlush(t *testing.T) {
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
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.King,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
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
					Rank:  poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
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
					Rank:  poker.Ace,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
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
			},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is.New(t).Equal(tc.expectedResult, poker.IsRoyalFlush(tc.cards))
		})
	}
}
