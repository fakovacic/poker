package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestCompareHands(t *testing.T) {
	cases := []struct {
		it             string
		hands          [][]*poker.Card
		expectedResult []int64
	}{
		{
			it: "compare high card and pair",
			hands: [][]*poker.Card{
				{
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
				{
					{
						Suite: poker.Hearts,
						Rank:  poker.Two,
					},
					{
						Suite: poker.Diamonds,
						Rank:  poker.Two,
					},
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
				},
			},
			expectedResult: []int64{
				1, 0,
			},
		},
		{
			it: "compare high card, high card and pair",
			hands: [][]*poker.Card{
				{
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
				{
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
						Rank:  poker.Five,
					},
				},
				{
					{
						Suite: poker.Hearts,
						Rank:  poker.Two,
					},
					{
						Suite: poker.Diamonds,
						Rank:  poker.Two,
					},
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
				},
			},
			expectedResult: []int64{
				2, 0, 1,
			},
		},
		{
			it: "compare pairs",
			hands: [][]*poker.Card{
				{
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
						Rank:  poker.Ten,
					},
					{
						Suite: poker.Spades,
						Rank:  poker.Ace,
					},
				},
				{
					{
						Suite: poker.Hearts,
						Rank:  poker.Ace,
					},
					{
						Suite: poker.Diamonds,
						Rank:  poker.Ace,
					},
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
				},
			},
			expectedResult: []int64{
				1, 0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			result := poker.CompareHands(tc.hands...)

			for i := range tc.expectedResult {
				if tc.expectedResult[i] != result[i] {
					t.Errorf("expected card %d: '%v' got: '%v'", i, tc.expectedResult[i], result[i])
				}
			}
		})
	}
}
