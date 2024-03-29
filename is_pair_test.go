package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestIsPair(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult bool
	}{
		{
			it: "ok, 2 cards",
			cards: []*poker.Card{
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
			it: "ok, 3 cards",
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
					Rank: poker.Queen,
				},
				{
					Rank: poker.Jack,
				},
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: true,
		},
		{
			it: "not ok, 2 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
				{
					Rank: poker.King,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 1 cards",
			cards: []*poker.Card{
				{
					Rank: poker.Ace,
				},
			},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.IsPair(tc.cards)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}

func benchmarkIsPair(cards []*poker.Card, b *testing.B) {
	for n := 0; n < b.N; n++ {
		poker.IsPair(cards)
	}
}

func BenchmarkIsPair(b *testing.B) {
	benchmarkIsPair([]*poker.Card{
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
			Rank: poker.Jack,
		},
		{
			Rank: poker.Ace,
		},
	}, b)
}
