package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestValidateSuite(t *testing.T) {
	cases := []struct {
		it             string
		suite          string
		expectedResult bool
	}{
		{
			it:             "clubs",
			suite:          poker.Clubs,
			expectedResult: true,
		},
		{
			it:             "diamonds",
			suite:          poker.Diamonds,
			expectedResult: true,
		},
		{
			it:             "hearts",
			suite:          poker.Hearts,
			expectedResult: true,
		},
		{
			it:             "spades",
			suite:          poker.Spades,
			expectedResult: true,
		},
		{
			it:             "ok",
			suite:          "ok",
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.ValidateSuite(tc.suite)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
