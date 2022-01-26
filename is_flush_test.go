package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
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
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is.New(t).Equal(tc.expectedResult, poker.IsFlush(tc.cards))
		})
	}
}
