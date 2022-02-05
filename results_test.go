package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestParseResult(t *testing.T) {
	cases := []struct {
		it             string
		result         string
		expectedOK     bool
		expectedResult poker.Result
	}{
		{
			it:             "high-card",
			result:         "high-card",
			expectedResult: poker.HighCard,
			expectedOK:     true,
		},
		{
			it:             "pair",
			result:         "pair",
			expectedResult: poker.Pair,
			expectedOK:     true,
		},
		{
			it:             "two-pairs",
			result:         "two-pairs",
			expectedResult: poker.TwoPairs,
			expectedOK:     true,
		},
		{
			it:             "three-of-a-kind",
			result:         "three-of-a-kind",
			expectedResult: poker.ThreeOfAKind,
			expectedOK:     true,
		},
		{
			it:             "straight",
			result:         "straight",
			expectedResult: poker.Straight,
			expectedOK:     true,
		},
		{
			it:             "flush",
			result:         "flush",
			expectedResult: poker.Flush,
			expectedOK:     true,
		},
		{
			it:             "full-house",
			result:         "full-house",
			expectedResult: poker.FullHouse,
			expectedOK:     true,
		},
		{
			it:             "four-of-a-kind",
			result:         "four-of-a-kind",
			expectedResult: poker.FourOfAKind,
			expectedOK:     true,
		},
		{
			it:             "straight-flush",
			result:         "straight-flush",
			expectedResult: poker.StraightFlush,
			expectedOK:     true,
		},
		{
			it:             "royal-flush",
			result:         "royal-flush",
			expectedResult: poker.RoyalFlush,
			expectedOK:     true,
		},
		{
			it:             "invalid suite",
			result:         "ok",
			expectedResult: "",
			expectedOK:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res, ok := poker.ParseResult(tc.result)
			if tc.expectedOK != ok {
				t.Errorf("expected ok: '%v' got: '%v'", tc.expectedOK, ok)
			}

			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
