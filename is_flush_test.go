package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestIsFlush(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult bool
	}{
		{
			it: "ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
			},
			expectedResult: true,
		},
		{
			it: "not ok, 5 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Diamonds,
				},
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
			},
			expectedResult: false,
		},
		{
			it: "not ok, 2 cards",
			cards: []*poker.Card{
				{
					Suite: poker.Clubs,
				},
				{
					Suite: poker.Clubs,
				},
			},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.IsFlush(tc.cards)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
